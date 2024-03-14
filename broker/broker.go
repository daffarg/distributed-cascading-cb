package broker

import (
	"context"
	"time"
)

type MessageBroker interface {
	Publish(ctx context.Context, topic string, message interface{}) error
	Subscribe(ctx context.Context, topic string) (string, error)

	// SubscribeAsync is used to subscribe to a topic and store the message with handler function
	SubscribeAsync(ctx context.Context, topic string, handler func(ctx context.Context, key, value string, exp time.Duration) error)
}
