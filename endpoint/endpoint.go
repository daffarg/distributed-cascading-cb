package endpoint

import (
	"context"

	"github.com/daffarg/distributed-cascading-cb/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

type CircuitBreakerEndpoint struct {
	GeneralEp endpoint.Endpoint
	GetEp     endpoint.Endpoint
	PostEp    endpoint.Endpoint
}

func NewCircuitBreakerEndpoint(svc service.CircuitBreakerService, log log.Logger) (CircuitBreakerEndpoint, error) {
	var generalEp endpoint.Endpoint
	{
		generalEp = makeGeneralEndpoint(svc)
	}

	var getEp endpoint.Endpoint
	{
		generalEp = makeGetEndpoint(svc)
	}

	var postEp endpoint.Endpoint
	{
		generalEp = makePostEndpoint(svc)
	}

	return CircuitBreakerEndpoint{
		GeneralEp: generalEp,
		GetEp:     getEp,
		PostEp:    postEp,
	}, nil
}

func (c *CircuitBreakerEndpoint) General(ctx context.Context, req *service.GeneralRequest) (*service.Response, error) {
	resp, err := c.GeneralEp(ctx, req)
	if err != nil {
		return &service.Response{}, err
	}

	return resp.(*service.Response), nil
}

func (c *CircuitBreakerEndpoint) Get(ctx context.Context, req *service.GetRequest) (*service.Response, error) {
	resp, err := c.GetEp(ctx, req)
	if err != nil {
		return &service.Response{}, err
	}

	return resp.(*service.Response), nil
}

func (c *CircuitBreakerEndpoint) Post(ctx context.Context, req *service.PostRequest) (*service.Response, error) {
	resp, err := c.PostEp(ctx, req)
	if err != nil {
		return &service.Response{}, err
	}

	return resp.(*service.Response), nil
}

func makeGeneralEndpoint(svc service.CircuitBreakerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*service.GeneralRequest)
		return svc.General(ctx, req)
	}
}

func makeGetEndpoint(svc service.CircuitBreakerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*service.GetRequest)
		return svc.Get(ctx, req)
	}
}

func makePostEndpoint(svc service.CircuitBreakerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*service.PostRequest)
		return svc.Post(ctx, req)
	}
}
