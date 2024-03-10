package transport

import (
	"github.com/daffarg/distributed-cascading-cb/protobuf"
	"github.com/go-kit/kit/endpoint"
)

type handler struct {
	request endpoint.Endpoint
	get     endpoint.Endpoint
	post    endpoint.Endpoint
	put     endpoint.Endpoint
	delete  endpoint.Endpoint
	protobuf.UnimplementedCircuitBreakerServer
}
