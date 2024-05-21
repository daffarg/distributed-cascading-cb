package service

import (
	"context"
	"github.com/daffarg/distributed-cascading-cb/util"
	"github.com/go-kit/log/level"
)

func (s *service) initConfig(ctx context.Context) {
	level.Info(s.log).Log(
		util.LogMessage, "circuit breaker config",
		util.LogConfig, s.config,
	)

	for _, ep := range s.config.AlternativeEndpoints {
		for _, alt := range ep.Alternatives {
			parsedUrl, _ := util.GetGeneralURLFormat(alt.Endpoint)
			endpointName := util.FormEndpointName(parsedUrl, alt.Method)
			_, err := s.repository.AddMembersIntoSet(
				ctx,
				util.FormRequiringEndpointsKey(endpointName),
				endpointName,
			)
			if err != nil {
				level.Error(s.log).Log(
					util.LogMessage, "failed to add requiring endpoint into set",
					util.LogError, err,
				)
			}
			encodedTopic := util.EncodeTopic(endpointName)
			go s.broker.SubscribeAsync(context.WithoutCancel(ctx), encodedTopic, s.repository.SetWithExp)
			s.subscribeMap[endpointName] = true
		}
	}
}
