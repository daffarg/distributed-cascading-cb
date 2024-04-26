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
		panic(err)
	}
	defer adminClient.Close()

	topicSpecification := []kafka.TopicSpecification{{
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 2,
	}}

	results, err := adminClient.CreateTopics(ctx, topicSpecification)
	if err != nil {
		return err
	}

	for _, result := range results {
		level.Info(k.log).Log(
			util.LogMessage, "result of kafka topic creation",
			util.LogTopic, result.Topic,
			util.LogErrorResult, result.Error,
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

func (k *kafkaBroker) SubscribeAsync(ctx context.Context, topic string, handler func(ctx context.Context, key, value string, exp time.Duration) error) {
	k.config["auto.offset.reset"] = "latest"
	k.config["enable.auto.commit"] = "false"

	consumer, _ := kafka.NewConsumer(&k.config)
	err := consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		level.Error(k.log).Log(
			util.LogMessage, "failed to subscribe to topic",
			util.LogTopic, topic,
			util.LogError, err,
		)
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

			err = handler(context.WithoutCancel(ctx), topic, msg.Status, time.Duration(msg.Timeout)*time.Second)
			if err != nil {
				level.Error(k.log).Log(
					util.LogMessage, "failed to store circuit breaker status into db",
					util.LogError, err,
					util.LogTopic, topic,
				)
				continue
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

//func (k *kafkaBroker) Publish(ctx context.Context, topic string, message *protobuf.Status) error {
//	conn, err := k.dialer.Dial(
//		"tcp",
//		k.config["bootstrap.servers"],
//	)
//	if err != nil {
//		return err
//	}
//	defer conn.Close()
//
//	controller, err := conn.Controller()
//	if err != nil {
//		return err
//	}
//	var controllerConn *kafka.Conn
//	controllerConn, err = k.dialer.Dial(
//		"tcp",
//		net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)),
//	)
//	if err != nil {
//		return err
//	}
//	defer controllerConn.Close()
//
//	topicConfigs := []kafka.TopicConfig{
//		{
//			Topic:             topic,
//			NumPartitions:     1,
//			ReplicationFactor: 3,
//		},
//	}
//
//	err = controllerConn.CreateTopics(topicConfigs...)
//	if err != nil {
//		return err
//	}
//
//	buf, err := proto.Marshal(message)
//	if err != nil {
//		return err
//	}
//
//	leaderConn, err := k.dialer.DialLeader(
//		ctx,
//		"tcp",
//		net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)),
//		topic,
//		0,
//	)
//	if err != nil {
//		return err
//	}
//
//	_, err = leaderConn.WriteMessages(kafka.Message{
//		Value: buf,
//	})
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (k *kafkaBroker) SubscribeAsync(ctx context.Context, topic string, handler func(ctx context.Context, key, value string, exp time.Duration) error) {
//	conn, err := k.dialer.Dial(
//		"tcp",
//		k.config["bootstrap.servers"],
//	)
//	if err != nil {
//		level.Error(k.log).Log(
//			util.LogMessage, "failed to dial kafka brokers",
//			util.LogError, err,
//			util.LogTopic, topic,
//		)
//	} else {
//		defer conn.Close()
//	}
//
//	controller, err := conn.Controller()
//	if err != nil {
//		level.Error(k.log).Log(
//			util.LogMessage, "failed to get kafka controller",
//			util.LogError, err,
//			util.LogTopic, topic,
//		)
//	}
//
//	leaderConn, err := k.dialer.DialLeader(
//		ctx,
//		"tcp",
//		net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)),
//		topic,
//		0,
//	)
//	if err != nil {
//		level.Error(k.log).Log(
//			util.LogMessage, "failed to dial kafka leader",
//			util.LogError, err,
//			util.LogTopic, topic,
//		)
//	} else {
//		defer leaderConn.Close()
//	}
//
//	batch := conn.ReadBatch(10, 1e6)
//	defer batch.Close()
//
//	b := make([]byte, 10e3) // 10KB max per message
//	for {
//		_, err := batch.Read(b)
//		if err != nil {
//			level.Error(k.log).Log(
//				util.LogMessage, "failed to read kafka message",
//				util.LogError, err,
//				util.LogTopic, topic,
//			)
//			continue
//		}
//
//		msg := &protobuf.Status{}
//		err = proto.Unmarshal(b, msg)
//		if err != nil {
//			level.Error(k.log).Log(
//				util.LogMessage, "failed to unmarshal kafka message",
//				util.LogError, err,
//				util.LogTopic, topic,
//			)
//			continue
//		}
//
//		err = handler(context.WithoutCancel(ctx), topic, msg.Status, time.Duration(msg.Timeout)*time.Second)
//		if err != nil {
//			level.Error(k.log).Log(
//				util.LogMessage, "failed to store circuit breaker status into db",
//				util.LogError, err,
//				util.LogTopic, topic,
//			)
//		}
//	}
//
//	//reader := kafka.NewReader(kafka.ReaderConfig{
//	//	Dialer:      k.dialer,
//	//	GroupID:     k.consumerGroup,
//	//	Brokers:     []string{k.config["bootstrap.servers"]},
//	//	Topic:       topic,
//	//	StartOffset: kafka.LastOffset,
//	//	Partition:   0,
//	//})
//	//
//	//level.Info(k.log).Log(
//	//	util.LogMessage, "subscribing to topic",
//	//	util.LogTopic, topic,
//	//)
//	//
//	//go func() {
//	//	for {
//	//		kafkaMsg, err := reader.FetchMessage(context.WithoutCancel(ctx))
//	//		if err != nil {
//	//			level.Error(k.log).Log(
//	//				util.LogMessage, "failed to read circuit breaker status from message broker",
//	//				util.LogError, err,
//	//				util.LogTopic, topic,
//	//			)
//	//		}
//	//
//	//		msg := &protobuf.Status{}
//	//		err = proto.Unmarshal(kafkaMsg.Value, msg)
//	//		if err != nil {
//	//			continue
//	//		}
//	//
//	//		err = handler(context.WithoutCancel(ctx), topic, msg.Status, time.Duration(msg.Timeout)*time.Second)
//	//		if err != nil {
//	//			level.Error(k.log).Log(
//	//				util.LogMessage, "failed to store circuit breaker status into db",
//	//				util.LogError, err,
//	//				util.LogTopic, topic,
//	//			)
//	//		}
//	//		err = reader.CommitMessages(ctx, kafkaMsg)
//	//		if err != nil {
//	//			level.Error(k.log).Log(
//	//				util.LogMessage, "failed to commit message",
//	//				util.LogError, err,
//	//				util.LogTopic, topic,
//	//			)
//	//		} else {
//	//			level.Info(k.log).Log(
//	//				util.LogMessage, "received and committed a message",
//	//				util.LogTopic, topic,
//	//				util.LogStatus, msg,
//	//			)
//	//		}
//	//
//	//	}
//	//}()
//}
