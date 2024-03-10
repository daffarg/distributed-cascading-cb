package broker

import "context"

type MessageBroker interface {
	Publish(ctx context.Context, topic string, message interface{}) error
	Subscribe(ctx context.Context, topic string, message chan<- string)
}
