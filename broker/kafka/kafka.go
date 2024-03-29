package kafka

import (
	"context"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/daffarg/distributed-cascading-cb/broker"
	"github.com/daffarg/distributed-cascading-cb/util"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/segmentio/kafka-go"
)

type kafkaBroker struct {
	log  log.Logger
	addr string
}

func NewKafkaBroker(log log.Logger, addr string) broker.MessageBroker {
	return &kafkaBroker{
		log:  log,
		addr: addr,
	}
}

func (k *kafkaBroker) Publish(ctx context.Context, topic string, message interface{}) error {
	writer := &kafka.Writer{
		Addr:  kafka.TCP(k.addr),
		Topic: topic,
	}

	var messageBytes []byte
	switch msg := message.(type) {
	case string:
		messageBytes = []byte(msg)
	case []byte:
		messageBytes = msg
	default:
		return util.ErrUnsupportedMessageType
	}

	return writer.WriteMessages(ctx, kafka.Message{
		Value: messageBytes,
	})
}

func (k *kafkaBroker) Subscribe(ctx context.Context, topic string) (string, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		GroupID:     os.Getenv("CB_CONSUMER_GROUP"),
		Brokers:     []string{k.addr},
		Topic:       topic,
		StartOffset: kafka.LastOffset,
	})

	msg, err := reader.ReadMessage(ctx)
	if err != nil {
		return "", err
	}

	return string(msg.Value), nil
}

func (k *kafkaBroker) SubscribeAsync(ctx context.Context, topic string, handler func(ctx context.Context, key, value string, exp time.Duration) error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		GroupID:     os.Getenv("CB_CONSUMER_GROUP"),
		Brokers:     []string{k.addr},
		Topic:       topic,
		StartOffset: kafka.LastOffset,
	})

	go func() {
		for {
			msg, err := reader.FetchMessage(ctx)
			if err != nil {
				level.Error(k.log).Log(
					util.LogMessage, "failed to read circuit breaker status from message broker",
					util.LogError, err,
					util.LogTopic, topic,
				)
			}

			if string(msg.Value) == "" {
				err := reader.CommitMessages(ctx, msg)
				if err != nil {
					level.Error(k.log).Log(
						util.LogMessage, "failed to commit message",
						util.LogError, err,
						util.LogTopic, topic,
					)
				}
				continue
			}

			messages := strings.Split(string(msg.Value), ":")

			cbStatus := messages[0]

			timeout, _ := strconv.ParseInt(messages[1], 10, 64)
			err = handler(ctx, topic, cbStatus, time.Duration(timeout)*time.Second)
			if err != nil {
				level.Error(k.log).Log(
					util.LogMessage, "failed to store circuit breaker status into db",
					util.LogError, err,
					util.LogTopic, topic,
				)
			}
			err = reader.CommitMessages(ctx, msg)
			if err != nil {
				level.Error(k.log).Log(
					util.LogMessage, "failed to commit message",
					util.LogError, err,
					util.LogTopic, topic,
				)
			}
		}
	}()
}
