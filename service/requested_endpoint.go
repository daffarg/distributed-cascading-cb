package service

import (
	"context"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/daffarg/distributed-cascading-cb/util"
	"github.com/go-kit/log/level"
)

type handleRequestedEndpointReq struct {
	RequestedEndpoint   string
	RequestedMethod     string
	CircuitBreakerName  string
	IsAlreadySubscribed bool
}

func (s *service) handleRequestedEndpoint(ctx context.Context, req *handleRequestedEndpointReq) {
	requiringEndpointsKey := util.FormRequiringEndpointsKey(req.CircuitBreakerName)

	_, err := s.repository.AddMembersIntoSet(context.WithoutCancel(ctx), requiringEndpointsKey, req.CircuitBreakerName)
	if err != nil {
		level.Error(s.log).Log(
			util.LogMessage, "failed to add requiring endpoint into set",
			util.LogError, err,
			util.LogRequest, req,
		)
	}

	if !req.IsAlreadySubscribed {
		encodedTopic := base58.Encode([]byte(req.CircuitBreakerName))
		go s.broker.SubscribeAsync(context.WithoutCancel(ctx), encodedTopic, s.repository.SetWithExp)
	}
}
