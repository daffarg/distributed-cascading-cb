package transport

import (
	"github.com/daffarg/distributed-cascading-cb/endpoint"
	"github.com/daffarg/distributed-cascading-cb/protobuf"
	"github.com/go-kit/kit/transport/grpc"
)

type handler struct {
	generalRequest grpc.Handler
	protobuf.UnimplementedCircuitBreakerServer
}

func NewCircuitBreakerServer(ep endpoint.CircuitBreakerEndpoint) protobuf.CircuitBreakerServer {
	opts := []grpc.ServerOption{}

	return &handler{
		generalRequest: grpc.NewServer(
			ep.GeneralRequestEp,
			decodeGeneralRequestReq,
			encodeResponse,
			opts...,
		),
	}
}
