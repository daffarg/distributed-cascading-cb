package transport

import (
	"context"

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

func (h *handler) GeneralRequest(ctx context.Context, req *protobuf.GeneralRequestInput) (*protobuf.Response, error) {
	_, res, err := h.generalRequest.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protobuf.Response), nil
}
