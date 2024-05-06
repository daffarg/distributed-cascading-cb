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
	PutEp     endpoint.Endpoint
	DeleteEp  endpoint.Endpoint
}

func NewCircuitBreakerEndpoint(svc service.CircuitBreakerService, log log.Logger) (CircuitBreakerEndpoint, error) {
	var generalEp endpoint.Endpoint
	{
		generalEp = makeGeneralEndpoint(svc)
	}

	var getEp endpoint.Endpoint
	{
		getEp = makeGetEndpoint(svc)
	}

	var postEp endpoint.Endpoint
	{
		postEp = makePostEndpoint(svc)
	}

	var putEp endpoint.Endpoint
	{
		putEp = makePutEndpoint(svc)
	}

	var deleteEp endpoint.Endpoint
	{
		deleteEp = makeDeleteEndpoint(svc)
	}

	return CircuitBreakerEndpoint{
		GeneralEp: generalEp,
		GetEp:     getEp,
		PostEp:    postEp,
		PutEp:     putEp,
		DeleteEp:  deleteEp,
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

func (c *CircuitBreakerEndpoint) Put(ctx context.Context, req *service.PutRequest) (*service.Response, error) {
	resp, err := c.PutEp(ctx, req)
	if err != nil {
		return &service.Response{}, err
	}

	return resp.(*service.Response), nil
}

func (c *CircuitBreakerEndpoint) Delete(ctx context.Context, req *service.DeleteRequest) (*service.Response, error) {
	resp, err := c.DeleteEp(ctx, req)
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

func makePutEndpoint(svc service.CircuitBreakerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*service.PutRequest)
		return svc.Put(ctx, req)
	}
}

func makeDeleteEndpoint(svc service.CircuitBreakerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*service.DeleteRequest)
		return svc.Delete(ctx, req)
	}
}
