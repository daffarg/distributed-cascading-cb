package service

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/daffarg/distributed-cascading-cb/circuitbreaker"
	"github.com/daffarg/distributed-cascading-cb/util"
	"github.com/go-kit/log/level"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GeneralRequestReq struct {
	Method            string            `json:"method" validate:"required"`
	URL               string            `json:"url" validate:"required"`
	Header            map[string]string `json:"header"`
	Body              []byte            `json:"body"`
	RequiringEndpoint string            `json:"requiring_endpoint"`
	RequiringMethod   string            `json:"requiring_method"`
}

func (s *service) GeneralRequest(ctx context.Context, req *GeneralRequestReq) (*Response, error) {
	if err := s.validator.Struct(req); err != nil {
		level.Error(s.log).Log(
			util.LogMessage, "failed precondition on request",
			util.LogError, err,
			util.LogRequest, req,
		)
		return &Response{}, status.Error(codes.FailedPrecondition, err.Error())
	}

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

	go func() { // add the endpoint itself
		if req.RequiringEndpoint != "" && req.RequiringMethod != "" {
			requiringEndpointsKey := util.FormRequiringEndpointsKey(circuitBreakerName)
			parsedUrl, err := util.GetGeneralURLFormat(req.RequiringEndpoint)
			if err != nil {
				level.Error(s.log).Log(
					util.LogMessage, "failed parsing requested url",
					util.LogError, err,
					util.LogRequest, req,
				)
			}

			requiringEndpointName := util.FormEndpointName(parsedUrl, req.RequiringMethod)
			isNewValue, err := s.repository.AddMembersIntoSet(context.WithoutCancel(ctx), requiringEndpointsKey, requiringEndpointName)
			if err != nil {
				level.Error(s.log).Log(
					util.LogMessage, "failed to add requiring endpoint into set",
					util.LogError, err,
					util.LogRequest, req,
				)
			}

			if isNewValue == 1 {
				go s.broker.SubscribeAsync(context.WithoutCancel(ctx), requiringEndpointName, s.repository.SetWithExp)
			}
		}
	}()

	go func() {
		requiringEndpointsKey := util.FormRequiringEndpointsKey(circuitBreakerName)

		isNewValue, err := s.repository.AddMembersIntoSet(context.WithoutCancel(ctx), requiringEndpointsKey, circuitBreakerName)
		if err != nil {
			level.Error(s.log).Log(
				util.LogMessage, "failed to add requiring endpoint into set",
				util.LogError, err,
				util.LogRequest, req,
			)
		}

		if isNewValue == 1 {
			go s.broker.SubscribeAsync(context.WithoutCancel(ctx), circuitBreakerName, s.repository.SetWithExp)
		}
	}()

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
		httpResponse, err := s.getCircuitBreaker(circuitBreakerName).Execute(func() (interface{}, error) {
			return s.doRequest(req.Method, req.URL, req.Body, req.Header)
		})
		if err != nil {
			if errors.Is(err, circuitbreaker.ErrOpenState) {
				return &Response{}, status.Error(codes.Unavailable, util.ErrCircuitBreakerOpen.Error())
			}
			level.Error(s.log).Log(
				util.LogMessage, "failed to execute the request",
				util.LogError, err,
				util.LogRequest, req,
			)
			return &Response{}, status.Error(codes.Internal, util.ErrFailedExecuteRequest.Error())
		}
		defer httpResponse.(*http.Response).Body.Close()

		response, err := s.convertToResponse(httpResponse.(*http.Response))
		if err != nil {
			level.Error(s.log).Log(
				util.LogMessage, "failed to convert http response to service response",
				util.LogError, err,
				util.LogRequest, req,
			)
		}

		return &response, nil
	}

	// cb is open
	return &Response{}, status.Error(codes.Unavailable, util.ErrCircuitBreakerOpen.Error())
}
