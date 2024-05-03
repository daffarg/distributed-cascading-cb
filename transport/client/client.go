package client

import (
	cbEndpoint "github.com/daffarg/distributed-cascading-cb/endpoint"
	"github.com/daffarg/distributed-cascading-cb/protobuf"
	"github.com/daffarg/distributed-cascading-cb/service"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

func NewGRPCClient(conn *grpc.ClientConn) service.CircuitBreakerService {
	var options []grpctransport.ClientOption

	var generalEndpoint endpoint.Endpoint
	{
		generalEndpoint = grpctransport.NewClient(
			conn,
			"protobuf.CircuitBreaker",
			"General",
			encodeGeneralRequest,
			decodeResponse,
			protobuf.Response{},
			options...,
		).Endpoint()
	}

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = grpctransport.NewClient(
			conn,
			"protobuf.CircuitBreaker",
			"Get",
			encodeGetRequest,
			decodeResponse,
			protobuf.Response{},
			options...,
		).Endpoint()
	}

	var postEndpoint endpoint.Endpoint
	{
		postEndpoint = grpctransport.NewClient(
			conn,
			"protobuf.CircuitBreaker",
			"Post",
			encodePostRequest,
			decodeResponse,
			protobuf.Response{},
			options...,
		).Endpoint()
	}

	return &cbEndpoint.CircuitBreakerEndpoint{
		GeneralEp: generalEndpoint,
		GetEp:     getEndpoint,
		PostEp:    postEndpoint,
	}
}
