package endpoint

import (
	"context"

	"github.com/daffarg/distributed-cascading-cb/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

type CircuitBreakerEndpoint struct {
	GeneralRequestEp endpoint.Endpoint
}

func NewCircuitBreakerEndpoint(svc service.CircuitBreakerService, log log.Logger) (CircuitBreakerEndpoint, error) {
	var generalRequestEp endpoint.Endpoint
	{
		generalRequestEp = makeGeneralRequestEndpoint(svc)
	}

	return CircuitBreakerEndpoint{
		GeneralRequestEp: generalRequestEp,
	}, nil
}

func makeGeneralRequestEndpoint(svc service.CircuitBreakerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(service.GeneralRequestReq)
		return svc.GeneralRequest(ctx, req)
	}
}
