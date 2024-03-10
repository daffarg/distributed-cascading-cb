package kafka

import (
	"context"

	"github.com/daffarg/distributed-cascading-cb/broker"
	"github.com/segmentio/kafka-go"
)

type kafkaBroker struct {
	addr string
}

func NewKafkaBroker(addr string) broker.MessageBroker {
	return &kafkaBroker{addr: addr}
}

func (k *kafkaBroker) Publish(ctx context.Context, topic string, message interface{}) error {
	writer := &kafka.Writer{
		Addr:  kafka.TCP(k.addr),
		Topic: topic,
	}

	return writer.WriteMessages(ctx, kafka.Message{
		Value: message.([]byte),
	})
}

func (k *kafkaBroker) Subscribe(ctx context.Context, topic string, message chan<- string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{k.addr},
		Topic:   topic,
	})

	go func() {
		for {
			msg, err := reader.ReadMessage(ctx)
			if err != nil {
				break
			}
			message <- string(msg.Value)
		}
	}()
}
