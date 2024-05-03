package service

import (
	"context"
	"errors"
	"github.com/daffarg/distributed-cascading-cb/config"
	"github.com/daffarg/distributed-cascading-cb/util"
	"github.com/go-kit/log/level"
)

type executeAlternativeEndpointReq struct {
	AlternativeEndpoint config.AlternativeEndpoint
	Body                []byte
	Header              map[string]string
}

func (s *service) executeAlternativeEndpoint(ctx context.Context, req *executeAlternativeEndpointReq) (*Response, error) {
	_, err := s.repository.Get(ctx, util.FormEndpointStatusKey(req.AlternativeEndpoint.AlternativeEndpoint))
	if err != nil {
		if !errors.Is(err, util.ErrKeyNotFound) {
			level.Error(s.log).Log(
				util.LogMessage, "failed to get cb alternative endpoint status from db",
				util.LogError, err,
				util.LogAlternativeEndpoint, req.AlternativeEndpoint,
			)
		}

		res, err := s.httpRequest(ctx, req.AlternativeEndpoint.Method, req.AlternativeEndpoint.Endpoint, req.Body, req.Header)
		if err != nil {
			level.Error(s.log).Log(
				util.LogMessage, "failed to execute the request to the alternative endpoint",
				util.LogError, err,
				util.LogAlternativeEndpoint, req.AlternativeEndpoint,
			)
			return nil, err
		}

		return res, nil
	}

	return &Response{}, err
}
