package endpoint

import "github.com/go-kit/kit/endpoint"

type CircuitBreakerEndpoint struct {
	Request endpoint.Endpoint
	Get     endpoint.Endpoint
	Post    endpoint.Endpoint
	Put     endpoint.Endpoint
	Delete  endpoint.Endpoint
}
