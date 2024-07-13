package broker

import (
	"context"
	"github.com/daffarg/distributed-cascading-cb/protobuf"
	"time"
)

type MessageBroker interface {
	Publish(ctx context.Context, topic string, message *protobuf.Status) error
	Subscribe(ctx context.Context, topic string) (*protobuf.Status, error)
	// SubscribeAsync is used to subscribe to a topic and store the message with handler function
	SubscribeAsync(request SubscribeAsyncRequest)
}

type SubscribeAsyncRequest struct {
	Ctx          context.Context
	Topic        string
	Set          func(ctx context.Context, key, value string, exp time.Duration) error
	Get          func(ctx context.Context, key string) (string, error)
	GetSetMember func(ctx context.Context, key string) ([]string, error)
}
