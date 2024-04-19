package service

import (
	"context"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/daffarg/distributed-cascading-cb/util"
	"github.com/go-kit/log/level"
)

type handleRequiringEndpointReq struct {
	RequiringEndpoint  string
	RequiringMethod    string
	CircuitBreakerName string
}

func (s *service) handleRequiringEndpoint(ctx context.Context, req *handleRequiringEndpointReq) {
	if req.RequiringEndpoint != "" && req.RequiringMethod != "" {
		requiringEndpointsKey := util.FormRequiringEndpointsKey(req.CircuitBreakerName)
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
			encodedTopic := base58.Encode([]byte(requiringEndpointName))
			go s.broker.SubscribeAsync(context.WithoutCancel(ctx), encodedTopic, s.repository.SetWithExp)
		}
	}
}
