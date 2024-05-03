package service

import (
	"context"
	"github.com/daffarg/distributed-cascading-cb/util"
	"github.com/go-kit/log/level"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GeneralRequest struct {
	Method            string            `json:"method" validate:"required"`
	URL               string            `json:"url" validate:"required"`
	Header            map[string]string `json:"header"`
	Body              []byte            `json:"body"`
	RequiringEndpoint string            `json:"requiring_endpoint" validate:"required"`
	RequiringMethod   string            `json:"requiring_method" validate:"required"`
}

func (s *service) General(ctx context.Context, req *GeneralRequest) (*Response, error) {
	if err := s.validator.Struct(req); err != nil {
		level.Error(s.log).Log(
			util.LogMessage, "failed precondition on request",
			util.LogError, err,
			util.LogRequest, req,
		)
		return &Response{}, status.Error(codes.FailedPrecondition, err.Error())
	}

	requestQuery := &request{
		Method:            req.Method,
		URL:               req.URL,
		Header:            req.Header,
		Body:              req.Body,
		RequiringEndpoint: req.RequiringEndpoint,
		RequiringMethod:   req.RequiringMethod,
	}

	return s.requestWithCircuitBreaker(ctx, requestQuery)
}
