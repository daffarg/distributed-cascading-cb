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

func (s *service) GeneralRequest(ctx context.Context, req GeneralRequestReq) (Response, error) {
	if err := s.validator.Struct(req); err != nil {
		level.Error(s.log).Log(
			util.LogMessage, "failed precondition on request",
			util.LogError, err,
			util.LogRequest, req,
		)
		return Response{}, status.Error(codes.FailedPrecondition, err.Error())
	}

	parsedUrl, err := util.GetGeneralURLFormat(req.URL)
	if err != nil {
		level.Error(s.log).Log(
			util.LogMessage, "failed parsing requested url",
			util.LogError, err,
			util.LogRequest, req,
		)
		return Response{}, status.Error(codes.Internal, util.ErrFailedParsingURL.Error())
	}

	req.URL = strings.ToLower(req.URL)
	req.Method = strings.ToUpper(req.Method)
	req.RequiringEndpoint = strings.ToLower(req.RequiringEndpoint)
	req.RequiringMethod = strings.ToUpper(req.RequiringMethod)

	circuitBreakerName := util.FormEndpointName(parsedUrl, req.Method)

	endpointStatusKey := util.FormEndpointStatusKey(circuitBreakerName)

	// check if we're already subscribing to the endpoint topics
	isMembers, err := s.repository.IsMembersOfSet(ctx, util.RequiredEndpointsKey, endpointStatusKey, circuitBreakerName)
	if err != nil {
		level.Error(s.log).Log(
			util.LogMessage, "failed to check if the endpoints has already subscribing to the topic",
			util.LogError, err,
			util.LogRequest, req,
		)
	}

	endpointMap := map[string]bool{
		endpointStatusKey:  isMembers[0],
		circuitBreakerName: isMembers[1],
	}

	for ep := range endpointMap {
		if !endpointMap[ep] {
			go s.broker.SubscribeAsync(context.WithoutCancel(ctx), ep, s.repository.SetWithExp)

			go func() {
				err = s.repository.AddMembersIntoSet(context.WithoutCancel(ctx), util.RequiredEndpointsKey, ep)
				if err != nil {
					level.Error(s.log).Log(
						util.LogMessage, "failed to add the endpoint into set",
						util.LogError, err,
						util.LogRequest, req,
					)
				}
			}()
		}
	}

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
				return Response{}, status.Error(codes.Unavailable, util.ErrCircuitBreakerOpen.Error())
			}
			return Response{}, status.Error(codes.Internal, util.ErrFailedExecuteRequest.Error())
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

		return response, nil
	}

	go func() {
		if req.RequiringEndpoint != "" && req.RequiringMethod != "" {
			requiringEndpointsKey := util.FormRequiringEndpointsKey(circuitBreakerName)

			requiringEndpointName := util.FormEndpointName(req.RequiringEndpoint, req.RequiringMethod)
			err := s.repository.AddMembersIntoSet(context.WithoutCancel(ctx), requiringEndpointsKey, requiringEndpointName)
			if err != nil {
				level.Error(s.log).Log(
					util.LogMessage, "failed to add requiring endpoint into set",
					util.LogError, err,
					util.LogRequest, req,
				)
			}
		}
	}()

	// cb is open
	return Response{}, status.Error(codes.Unavailable, util.ErrCircuitBreakerOpen.Error())
}
