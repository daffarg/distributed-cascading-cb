package service

import (
	"context"
	"errors"
	"github.com/daffarg/distributed-cascading-cb/circuitbreaker"
	"github.com/daffarg/distributed-cascading-cb/util"
	"github.com/go-kit/log/level"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

type request struct {
	Method            string            `json:"method"`
	URL               string            `json:"url"`
	Header            map[string]string `json:"header"`
	Body              []byte            `json:"body"`
	RequiringEndpoint string            `json:"requiring_endpoint"`
	RequiringMethod   string            `json:"requiring_method"`
}

func (s *service) requestWithCircuitBreaker(ctx context.Context, req *request) (*Response, error) {
	parsedUrl, err := util.GetGeneralURLFormat(req.URL)
	if err != nil {
		level.Error(s.log).Log(
			util.LogMessage, "failed parsing requested url",
			util.LogError, err,
			util.LogRequest, req,
		)
		return &Response{}, status.Error(codes.Internal, util.ErrFailedParsingURL.Error())
	}

	req.URL = strings.ToLower(req.URL)
	req.Method = strings.ToUpper(req.Method)
	req.RequiringEndpoint = strings.ToLower(req.RequiringEndpoint)
	req.RequiringMethod = strings.ToUpper(req.RequiringMethod)

	circuitBreakerName := util.FormEndpointName(parsedUrl, req.Method)
	endpointStatusKey := util.FormEndpointStatusKey(circuitBreakerName)

	topic := util.EncodeTopic(circuitBreakerName)
	msg, err := s.broker.Subscribe(ctx, topic)
	if err != nil {
		level.Error(s.log).Log(
			util.LogMessage, "failed to subscribe to topic",
			util.LogError, err,
			util.LogRequest, req,
			util.LogTopic, topic,
		)
		return &Response{}, status.Error(codes.Internal, util.ErrFailedSubscribe.Error())
	}

	if msg.Status == circuitbreaker.StateOpen.String() {
		err = s.repository.SetWithExp(ctx, util.FormEndpointStatusKey(msg.Endpoint), msg.Status, time.Duration(msg.Timeout)*time.Second)
		if err != nil {
			level.Error(s.log).Log(
				util.LogMessage, "failed to store circuit breaker status into db",
				util.LogError, err,
				util.LogRequest, req,
				util.LogTopic, topic,
				util.LogStatus, msg,
			)
		}

		return &Response{}, status.Error(codes.Unavailable, util.ErrCircuitBreakerOpen.Error())
	}

	go s.handleRequiringEndpoint(ctx, &handleRequiringEndpointReq{
		RequiringEndpoint:  req.RequiringEndpoint,
		RequiringMethod:    req.RequiringMethod,
		CircuitBreakerName: circuitBreakerName,
	})

	isAlreadySubscribed := false
	_, ok := s.breakers[circuitBreakerName]
	if ok {
		isAlreadySubscribed = true
	}

	go s.handleRequestedEndpoint(ctx, &handleRequestedEndpointReq{
		RequestedEndpoint:   req.URL,
		RequestedMethod:     req.Method,
		CircuitBreakerName:  circuitBreakerName,
		IsAlreadySubscribed: isAlreadySubscribed,
	})

	altEndpoint, hasAltEp := s.config.AlternativeEndpoints[circuitBreakerName]

	_, err = s.repository.Get(ctx, endpointStatusKey)
	if err != nil {
		if !errors.Is(err, util.ErrKeyNotFound) {
			level.Error(s.log).Log(
				util.LogMessage, "failed to get cb status from db",
				util.LogError, err,
				util.LogRequest, req,
			)
		}

		// do request if error when getting cb status or cb status is not open
		response, err := s.getCircuitBreaker(circuitBreakerName).Execute(func() (interface{}, error) {
			return s.httpRequest(ctx, req.Method, req.URL, req.Body, req.Header)
		})
		if err != nil {
			if errors.Is(err, circuitbreaker.ErrOpenState) {
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
				}
			}
			return &Response{}, status.Error(codes.Internal, util.ErrFailedExecuteRequest.Error())
		}

		return response.(*Response), nil
	}

	// cb is open
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
	}

	return &Response{}, status.Error(codes.Unavailable, util.ErrCircuitBreakerOpen.Error())
}
