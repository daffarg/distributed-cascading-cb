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

	var generalRequestEndpoint endpoint.Endpoint
	{
		generalRequestEndpoint = grpctransport.NewClient(
			conn,
			"pb.CircuitBreaker",
			"GeneralRequest",
			encodeGeneralRequestReq,
			decodeResponse,
			protobuf.Response{},
			options...,
		).Endpoint()
	}

	return &cbEndpoint.CircuitBreakerEndpoint{
		GeneralRequestEp: generalRequestEndpoint,
	}
}
