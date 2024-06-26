package transport

import (
	"context"

	"github.com/daffarg/distributed-cascading-cb/endpoint"
	"github.com/daffarg/distributed-cascading-cb/protobuf"
	"github.com/go-kit/kit/transport/grpc"
)

type handler struct {
	general grpc.Handler
	get     grpc.Handler
	post    grpc.Handler
	put     grpc.Handler
	delete  grpc.Handler
	protobuf.UnimplementedCircuitBreakerServer
}

func NewCircuitBreakerServer(ep endpoint.CircuitBreakerEndpoint) protobuf.CircuitBreakerServer {
	opts := []grpc.ServerOption{}

	return &handler{
		general: grpc.NewServer(
			ep.GeneralEp,
			decodeGeneralRequest,
			encodeResponse,
			opts...,
		),
		get: grpc.NewServer(
			ep.GetEp,
			decodeGetRequest,
			encodeResponse,
			opts...,
		),
		post: grpc.NewServer(
			ep.PostEp,
			decodePostRequest,
			encodeResponse,
			opts...,
		),
		put: grpc.NewServer(
			ep.PostEp,
			decodePutRequest,
			encodeResponse,
			opts...,
		),
		delete: grpc.NewServer(
			ep.PostEp,
			decodeDeleteRequest,
			encodeResponse,
			opts...,
		),
	}
}

func (h *handler) General(ctx context.Context, req *protobuf.GeneralRequest) (*protobuf.Response, error) {
	_, res, err := h.general.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protobuf.Response), nil
}

func (h *handler) Get(ctx context.Context, req *protobuf.GetRequest) (*protobuf.Response, error) {
	_, res, err := h.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protobuf.Response), nil
}

func (h *handler) Post(ctx context.Context, req *protobuf.PostRequest) (*protobuf.Response, error) {
	_, res, err := h.post.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protobuf.Response), nil
}

func (h *handler) Put(ctx context.Context, req *protobuf.PutRequest) (*protobuf.Response, error) {
	_, res, err := h.put.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protobuf.Response), nil
}

func (h *handler) Delete(ctx context.Context, req *protobuf.DeleteRequest) (*protobuf.Response, error) {
	_, res, err := h.delete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*protobuf.Response), nil
}
