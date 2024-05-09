package kafka

import (
	"bufio"
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/daffarg/distributed-cascading-cb/broker"
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
	config kafka.ConfigMap
	log    log.Logger
}

func NewKafkaBroker(log log.Logger, configPath string) (broker.MessageBroker, error) {
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
		config: m,
		log:    log,
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

	if len(metadata.Topics[topic].Partitions) <= 0 {
		return nil, util.ErrUpdatedStatusNotFound
	}

	consumerConfig := k.config
	consumerConfig["default.topic.config"] = kafka.ConfigMap{"auto.offset.reset": "earliest"}
	consumerConfig["enable.auto.commit"] = "false"
	consumerConfig["allow.auto.create.topics"] = "true"

	consumer, err := kafka.NewConsumer(&consumerConfig)
	if err != nil {
		return nil, err
	}
	defer consumer.Close()

	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(util.GetIntEnv("FIRST_SUBSCRIBE_TIMEOUT", 100))*time.Millisecond)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, util.ErrUpdatedStatusNotFound
		default:
			ev := consumer.Poll(util.GetIntEnv("FIRST_SUBSCRIBE_TIMEOUT", 100))
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case kafka.AssignedPartitions:
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

			}
		}
	}
}

func (k *kafkaBroker) SubscribeAsync(ctx context.Context, topic string, handler func(ctx context.Context, key, value string, exp time.Duration) error) {
	adminClient, err := kafka.NewAdminClient(&k.config)
	if err != nil {
		level.Error(k.log).Log(
			util.LogMessage, "failed to get kafka admin client",
			util.LogTopic, topic,
			util.LogError, err,
		)
	} else {
		metadata, err := adminClient.GetMetadata(&topic, false, util.GetIntEnv("KAFKA_GET_METADATA_TIMEOUT", 30000))
		if err != nil {
			level.Error(k.log).Log(
				util.LogMessage, "failed to get metadata of kafka topic",
				util.LogTopic, topic,
				util.LogError, err,
			)
		}

		level.Info(k.log).Log(
			util.LogMessage, "metadata of kafka topic",
			util.LogTopic, topic,
			util.LogMetadata, metadata,
		)

		if len(metadata.Topics[topic].Partitions) <= 0 { // determine if the topic not exists
			level.Info(k.log).Log(
				util.LogMessage, "kafka topic not found, will be creating the topic",
				util.LogTopic, topic,
			)

			topicSpecification := []kafka.TopicSpecification{{
				Topic:             topic,
				NumPartitions:     1,
				ReplicationFactor: 3,
			}}

			results, err := adminClient.CreateTopics(ctx, topicSpecification)
			if err != nil {
				level.Error(k.log).Log(
					util.LogMessage, "failed to create a new kafka topic",
					util.LogTopic, topic,
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
		err := consumer.SubscribeTopics([]string{topic}, nil)
		if err != nil {
			level.Warn(k.log).Log(
				util.LogMessage, "failed to subscribe to topic",
				util.LogTopic, topic,
				util.LogError, err,
			)
			time.Sleep(time.Duration(util.GetIntEnv("RETRY_SUBSCRIBE_INTERVAL", 10)) * time.Second)
		} else {
			break
		}
	}
	defer consumer.Close()

	for {
		kafkaMsg, err := consumer.ReadMessage(-1)
		if err != nil {
			level.Error(k.log).Log(
				util.LogMessage, "failed to read a message from kafka",
				util.LogTopic, topic,
				util.LogError, err,
			)
		} else {
			msg := &protobuf.Status{}
			err = proto.Unmarshal(kafkaMsg.Value, msg)
			if err != nil {
				level.Error(k.log).Log(
					util.LogMessage, "failed to unmarshal kafka message",
					util.LogError, err,
					util.LogTopic, topic,
				)
				continue
			}

			timestamp, _ := time.Parse(time.RFC3339, msg.Timestamp)
			expiredTime := timestamp.Add(time.Duration(msg.Timeout) * time.Second)
			if time.Now().Before(expiredTime) {
				timeout := expiredTime.Sub(time.Now()) * time.Second
				err = handler(context.WithoutCancel(ctx), util.FormEndpointStatusKey(msg.Endpoint), msg.Status, timeout)
				if err != nil {
					level.Error(k.log).Log(
						util.LogMessage, "failed to store circuit breaker status into db",
						util.LogError, err,
						util.LogTopic, topic,
						util.LogStatus, msg,
					)
					continue
				}
			}

			_, err = consumer.CommitMessage(kafkaMsg)
			if err != nil {
				level.Error(k.log).Log(
					util.LogMessage, "failed to commit a kafka message",
					util.LogError, err,
					util.LogTopic, topic,
				)
			} else {
				level.Info(k.log).Log(
					util.LogMessage, "received and committed a kafka message",
					util.LogTopic, topic,
					util.LogStatus, msg,
				)
			}
		}
	}
}
