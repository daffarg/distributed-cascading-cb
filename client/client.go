package client

import (
	"context"

	"github.com/daffarg/distributed-cascading-cb/protobuf"
	"google.golang.org/grpc"
)

type CircuitBreakerClient struct {
	client protobuf.CircuitBreakerClient
}

func NewCircuitBreakerClient(conn *grpc.ClientConn) *CircuitBreakerClient {
	return &CircuitBreakerClient{
		client: protobuf.NewCircuitBreakerClient(conn),
	}
}

func (c *CircuitBreakerClient) GeneralRequest(ctx context.Context, req *protobuf.GeneralRequestInput) (*protobuf.Response, error) {

	res, err := c.client.GeneralRequest(ctx, req)
	if err != nil {
		return &protobuf.Response{}, err
	}

	return res, nil
}
