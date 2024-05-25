package service

import (
	"context"
	"github.com/daffarg/distributed-cascading-cb/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) handleCircuitBreakerOpen(ctx context.Context, circuitBreakerName string, req *request) (*Response, error) {
	altEndpoint, hasAltEp := s.config.AlternativeEndpoints[circuitBreakerName]

	isException := false
	_, ok := s.config.Exceptions[circuitBreakerName]
	if ok {
		isException = true
	}

	if hasAltEp {
		res, err := s.executeAlternativeEndpoint(ctx, &executeAlternativeEndpointReq{
			AlternativeEndpoint: altEndpoint,
			Body:                req.Body,
			Header:              req.Header,
		})
		if err != nil {
			return &Response{}, status.Error(codes.Internal, util.ErrFailedExecuteAltEndpoint.Error())
		}
		res.IsFromAlternativeEndpoint = true
		return res, nil
	} else if isException {
		res, err := s.httpRequest(ctx, req.Method, req.URL, req.Body, req.Header)
		if err != nil {
			return &Response{}, status.Error(codes.Internal, util.ErrFailedExecuteRequest.Error())
		}
		return res, nil
	} else {
		return &Response{}, status.Error(codes.Unavailable, util.ErrCircuitBreakerOpen.Error())
	}
}
