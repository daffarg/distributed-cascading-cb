package service

import (
	"context"
	"errors"
	"github.com/daffarg/distributed-cascading-cb/config"
	"github.com/daffarg/distributed-cascading-cb/util"
	"github.com/go-kit/log/level"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type executeAlternativeEndpointReq struct {
	AlternativeEndpoint config.AlternativeEndpoint
	Body                []byte
	Header              map[string]string
}

func (s *service) executeAlternativeEndpoint(ctx context.Context, req *executeAlternativeEndpointReq) (*Response, error) {
	for _, alt := range req.AlternativeEndpoint.Alternatives {
		endpoint := util.FormEndpointName(alt.Endpoint, alt.Method)
		_, err := s.repository.Get(ctx, util.FormEndpointStatusKey(endpoint))
		if err != nil {
			if !errors.Is(err, util.ErrKeyNotFound) {
				level.Error(s.log).Log(
					util.LogMessage, "failed to get cb alternative endpoint status from db",
					util.LogError, err,
					util.LogAlternativeEndpoint, endpoint,
				)
			}

			// do request if error when getting cb status or cb status is not open
			response, err := s.getCircuitBreaker(endpoint).Execute(func() (interface{}, error) {
				return s.httpRequest(ctx, alt.Method, alt.Endpoint, req.Body, req.Header)
			})
			if err != nil {
				level.Error(s.log).Log(
					util.LogMessage, "failed to execute the request to the alternative endpoint",
					util.LogError, err,
					util.LogAlternativeEndpoint, endpoint,
				)
			} else {
				return response.(*Response), nil
			}
		}
	}

	// If all alternative endpoints are in open state or failed to execute the requests
	return &Response{}, status.Error(codes.Internal, util.ErrFailedExecuteAltEndpoint.Error())
}
