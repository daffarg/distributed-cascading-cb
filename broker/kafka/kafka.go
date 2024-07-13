package kafka

import (
	"bufio"
	"context"
	"errors"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/daffarg/distributed-cascading-cb/broker"
	"github.com/daffarg/distributed-cascading-cb/config"
	"github.com/daffarg/distributed-cascading-cb/protobuf"
	"github.com/daffarg/distributed-cascading-cb/util"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"google.golang.org/protobuf/proto"
	"os"
	"strings"
	"time"
)

type kafkaBroker struct {
	config   kafka.ConfigMap
	log      log.Logger
	cbConfig *config.Config
}

func NewKafkaBroker(log log.Logger, configPath string, cbConfig *config.Config) (broker.MessageBroker, error) {
	m := make(map[string]kafka.ConfigValue)

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			kv := strings.Split(line, "=")
			parameter := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			m[parameter] = value
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	m["group.id"] = os.Getenv("CB_CONSUMER_GROUP")

	return &kafkaBroker{
		config:   m,
		log:      log,
		cbConfig: cbConfig,
	}, nil
}

func (k *kafkaBroker) Publish(ctx context.Context, topic string, message *protobuf.Status) error {
	adminClient, err := kafka.NewAdminClient(&k.config)
	if err != nil {
		return err
	}
	defer adminClient.Close()

	topicSpecification := []kafka.TopicSpecification{{
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 3,
	}}

	results, err := adminClient.CreateTopics(ctx, topicSpecification)
	if err != nil {
		return err
	}

	for _, result := range results {
		level.Info(k.log).Log(
			util.LogMessage, "result of kafka topic creation",
			util.LogTopic, result.Topic,
			util.LogResult, result.Error,
		)
	}

	p, _ := kafka.NewProducer(&k.config)

	msgBuf, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msgBuf,
	}, nil)
	if err != nil {
		return err
	}

	p.Flush(15 * 1000)
	p.Close()

	return nil
}

func (k *kafkaBroker) Subscribe(_ context.Context, topic string) (*protobuf.Status, error) {
	adminClient, err := kafka.NewAdminClient(&k.config)
	if err != nil {
		return nil, err
	}
	defer adminClient.Close()

	metadata, err := adminClient.GetMetadata(&topic, false, util.GetIntEnv("KAFKA_GET_METADATA_TIMEOUT", 30000))
	if err != nil {
		return nil, err
	}

	level.Info(k.log).Log(
		util.LogMessage, "metadata of kafka topic",
		util.LogTopic, topic,
		util.LogMetadata, metadata,
	)

	if len(metadata.Topics[topic].Partitions) <= 0 {
		return nil, util.ErrUpdatedStatusNotFound
	}

	consumerConfig := make(kafka.ConfigMap)
	for k, v := range k.config {
		consumerConfig[k] = v
	}

	consumerConfig["default.topic.config"] = kafka.ConfigMap{"auto.offset.reset": "earliest"}
	consumerConfig["enable.auto.commit"] = false
	consumerConfig["go.application.rebalance.enable"] = true

	consumer, err := kafka.NewConsumer(&consumerConfig)
	if err != nil {
		return nil, err
	}
	defer consumer.Close()

	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(util.GetIntEnv("GET_STATUS_FIRST_TIME_TIMEOUT", 100))*time.Millisecond)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			level.Info(k.log).Log(
				util.LogMessage, "timeout when polling a message from kafka",
				util.LogTopic, topic,
			)
			return nil, util.ErrUpdatedStatusNotFound
		default:
			ev := consumer.Poll(util.GetIntEnv("FIRST_POLL_TIMEOUT", 100))
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				level.Info(k.log).Log(
					util.LogMessage, "assigning kafka partitions",
					util.LogTopic, topic,
				)

				parts := make([]kafka.TopicPartition, len(e.Partitions))
				for i, tp := range e.Partitions {
					tp.Offset = kafka.OffsetTail(1)
					parts[i] = tp
				}

				err := consumer.Assign(parts)
				if err != nil {
					level.Error(k.log).Log(
						util.LogMessage, "failed to assign partitions",
						util.LogError, err,
						util.LogTopic, topic,
					)
				}
			case *kafka.Message:
				msg := &protobuf.Status{}
				err = proto.Unmarshal(e.Value, msg)
				if err != nil {
					return nil, err
				}

				level.Info(k.log).Log(
					util.LogMessage, "received a new cb status from kafka",
					util.LogTopic, topic,
					util.LogStatus, msg,
				)

				_, err = consumer.CommitMessage(e)
				if err != nil {
					level.Error(k.log).Log(
						util.LogMessage, "failed to commit a status message to kafka",
						util.LogError, err,
						util.LogTopic, topic,
						util.LogStatus, msg,
					)
				} else {
					level.Info(k.log).Log(
						util.LogMessage, "received and committed a kafka message",
						util.LogTopic, topic,
						util.LogStatus, msg,
					)
				}

				return msg, nil
			case kafka.PartitionEOF:
				level.Info(k.log).Log(
					util.LogMessage, "reached the end of a partition",
					util.LogTopic, topic,
					util.LogEvent, e,
				)
			case kafka.Error:
				level.Error(k.log).Log(
					util.LogMessage, "error when polling a message from kafka",
					util.LogError, err,
					util.LogTopic, topic,
				)
				return nil, err
			default:
				level.Info(k.log).Log(
					util.LogMessage, "default case when polling a message from kafka",
					util.LogTopic, topic,
					util.LogEvent, e,
				)
			}
		}
	}
}

func (k *kafkaBroker) SubscribeAsync(request broker.SubscribeAsyncRequest) {
	adminClient, err := kafka.NewAdminClient(&k.config)
	if err != nil {
		level.Error(k.log).Log(
			util.LogMessage, "failed to get kafka admin client",
			util.LogTopic, request.Topic,
			util.LogError, err,
		)
	} else {
		metadata, err := adminClient.GetMetadata(&request.Topic, false, util.GetIntEnv("KAFKA_GET_METADATA_TIMEOUT", 30000))
		if err != nil {
			level.Error(k.log).Log(
				util.LogMessage, "failed to get metadata of kafka topic",
				util.LogTopic, request.Topic,
				util.LogError, err,
			)
		}

		level.Info(k.log).Log(
			util.LogMessage, "metadata of kafka topic",
			util.LogTopic, request.Topic,
			util.LogMetadata, metadata,
		)

		if len(metadata.Topics[request.Topic].Partitions) <= 0 { // determine if the topic not exists
			level.Info(k.log).Log(
				util.LogMessage, "kafka topic not found, will be creating the topic",
				util.LogTopic, request.Topic,
			)

			topicSpecification := []kafka.TopicSpecification{{
				Topic:             request.Topic,
				NumPartitions:     1,
				ReplicationFactor: 3,
			}}

			results, err := adminClient.CreateTopics(request.Ctx, topicSpecification)
			if err != nil {
				level.Error(k.log).Log(
					util.LogMessage, "failed to create a new kafka topic",
					util.LogTopic, request.Topic,
					util.LogError, err,
				)
			}

			for _, result := range results {
				level.Info(k.log).Log(
					util.LogMessage, "result of kafka topic creation",
					util.LogTopic, result.Topic,
					util.LogResult, result.Error,
				)
			}
		}

		adminClient.Close()
	}

	k.config["auto.offset.reset"] = "latest"
	k.config["enable.auto.commit"] = "false"
	k.config["allow.auto.create.topics"] = "true"

	consumer, _ := kafka.NewConsumer(&k.config)

	for {
		err := consumer.SubscribeTopics([]string{request.Topic}, nil)
		if err != nil {
			level.Warn(k.log).Log(
				util.LogMessage, "failed to subscribe to topic",
				util.LogTopic, request.Topic,
				util.LogError, err,
			)
			time.Sleep(time.Duration(util.GetIntEnv("RETRY_SUBSCRIBE_INTERVAL", 10)) * time.Second)
		} else {
			break
		}
	}
	defer consumer.Close()

	level.Info(k.log).Log(
		util.LogMessage, "subscribed to a kafka topic",
		util.LogTopic, request.Topic,
	)

	for {
		kafkaMsg, err := consumer.ReadMessage(-1)
		if err != nil {
			level.Error(k.log).Log(
				util.LogMessage, "failed to read a message from kafka",
				util.LogTopic, request.Topic,
				util.LogError, err,
			)
		} else {
			msg := &protobuf.Status{}
			err = proto.Unmarshal(kafkaMsg.Value, msg)
			if err != nil {
				level.Error(k.log).Log(
					util.LogMessage, "failed to unmarshal kafka message",
					util.LogError, err,
					util.LogTopic, request.Topic,
				)
				continue
			}

			_, err = consumer.CommitMessage(kafkaMsg)
			if err != nil {
				level.Error(k.log).Log(
					util.LogMessage, "failed to commit a kafka message",
					util.LogError, err,
					util.LogTopic, request.Topic,
				)
			} else {
				level.Info(k.log).Log(
					util.LogMessage, "received and committed a kafka message",
					util.LogTopic, request.Topic,
					util.LogStatus, msg,
				)
			}

			timestamp, _ := time.Parse(time.RFC3339, msg.Timestamp)
			expiredTime := timestamp.Add(time.Duration(msg.Timeout) * time.Second)
			if time.Now().Before(expiredTime) {
				timeout := expiredTime.Sub(time.Now()) * time.Second
				isThereAlt := false
				if alt, ok := k.cbConfig.AlternativeEndpoints[msg.Endpoint]; ok {
					for _, ep := range alt.Alternatives {
						endpointName := util.FormEndpointName(ep.Endpoint, ep.Method)
						_, err := request.Get(context.Background(), util.FormEndpointStatusKey(endpointName))
						if err != nil {
							if errors.Is(err, util.ErrKeyNotFound) {
								isThereAlt = true
								break
							} else {
								level.Error(k.log).Log(
									util.LogMessage, "failed to get endpoint status from db",
									util.LogEndpoint, endpointName,
									util.LogError, err,
								)
							}
						}
					}
				}

				if !isThereAlt {
					go func() {
						requiringEndpoints, err := request.GetSetMember(context.Background(), util.FormRequiringEndpointsKey(msg.Endpoint))
						if err != nil {
							level.Error(k.log).Log(
								util.LogMessage, "failed to get requiring endpoints from db",
								util.LogError, err,
								util.LogCircuitBreakerEndpoint, msg.Endpoint,
								util.LogCircuitBreakerNewStatus, msg.Status,
							)

							err := request.Set(context.Background(), util.FormEndpointStatusKey(msg.Endpoint), msg.Status, timeout)
							if err != nil {
								level.Error(k.log).Log(
									util.LogMessage, "failed to set circuit breaker status to db",
									util.LogError, err,
									util.LogCircuitBreakerEndpoint, msg.Endpoint,
									util.LogCircuitBreakerNewStatus, msg.Status,
								)
							}
						} else {
							for _, ep := range requiringEndpoints {
								encodedTopic := util.EncodeTopic(ep)
								message := &protobuf.Status{
									Endpoint:  ep,
									Status:    msg.Status,
									Timeout:   uint32(util.GetIntEnv("CB_TIMEOUT", 60)),
									Timestamp: time.Now().Format(time.RFC3339),
								}
								if err != nil {
									level.Error(k.log).Log(
										util.LogMessage, "failed to marshal circuit breaker status",
										util.LogError, err,
										util.LogCircuitBreakerEndpoint, ep,
										util.LogCircuitBreakerNewStatus, msg.Status,
									)
								}

								err = k.Publish(context.Background(), encodedTopic, message)
								if err != nil {
									level.Error(k.log).Log(
										util.LogMessage, "failed to publish circuit breaker status",
										util.LogError, err,
										util.LogCircuitBreakerEndpoint, ep,
										util.LogCircuitBreakerNewStatus, msg.Status,
									)
								} else {
									level.Info(k.log).Log(
										util.LogMessage, "published circuit breaker status",
										util.LogStatus, message,
									)
								}

								err := request.Set(context.Background(), util.FormEndpointStatusKey(ep), msg.Status, timeout)
								if err != nil {
									level.Error(k.log).Log(
										util.LogMessage, "failed to set circuit breaker status to db",
										util.LogError, err,
										util.LogCircuitBreakerEndpoint, ep,
										util.LogCircuitBreakerNewStatus, msg.Status,
									)
								}
							}
						}
					}()
				} else {
					level.Info(k.log).Log(
						util.LogMessage, "there are still alternative endpoints, only publishing the endpoint not its requirings",
						util.LogCircuitBreakerEndpoint, msg.Endpoint,
						util.LogCircuitBreakerNewStatus, msg.Status,
					)

					err := request.Set(context.Background(), util.FormEndpointStatusKey(msg.Endpoint), msg.Status, timeout)
					if err != nil {
						level.Error(k.log).Log(
							util.LogMessage, "failed to set circuit breaker status to db",
							util.LogError, err,
							util.LogCircuitBreakerEndpoint, msg.Endpoint,
							util.LogCircuitBreakerNewStatus, msg.Status,
						)
					}

					encodedTopic := util.EncodeTopic(msg.Endpoint)
					message := &protobuf.Status{
						Endpoint:  msg.Endpoint,
						Status:    msg.Status,
						Timeout:   uint32(util.GetIntEnv("CB_TIMEOUT", 60)),
						Timestamp: time.Now().Format(time.RFC3339),
					}
					if err != nil {
						level.Error(k.log).Log(
							util.LogMessage, "failed to marshal circuit breaker status",
							util.LogError, err,
							util.LogCircuitBreakerEndpoint, msg.Endpoint,
							util.LogCircuitBreakerNewStatus, msg.Status,
						)
					}

					err = k.Publish(context.Background(), encodedTopic, message)
					if err != nil {
						level.Error(k.log).Log(
							util.LogMessage, "failed to publish circuit breaker status",
							util.LogError, err,
							util.LogCircuitBreakerEndpoint, msg.Endpoint,
							util.LogCircuitBreakerNewStatus, msg.Status,
						)
					} else {
						level.Info(k.log).Log(
							util.LogMessage, "published circuit breaker status",
							util.LogStatus, message,
						)
					}
				}
			}

			_, err = consumer.CommitMessage(kafkaMsg)
			if err != nil {
				level.Error(k.log).Log(
					util.LogMessage, "failed to commit a kafka message",
					util.LogError, err,
					util.LogTopic, request.Topic,
				)
			} else {
				level.Info(k.log).Log(
					util.LogMessage, "received and committed a kafka message",
					util.LogTopic, request.Topic,
					util.LogStatus, msg,
				)
			}
		}
	}
}
