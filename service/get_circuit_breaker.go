package service

import (
	"context"
	"time"

	"github.com/daffarg/distributed-cascading-cb/circuitbreaker"
	"github.com/daffarg/distributed-cascading-cb/util"
	"github.com/go-kit/log/level"
)

func (s *service) getCircuitBreaker(name string) *circuitbreaker.CircuitBreaker {
	if cb, ok := s.breakers[name]; ok {
		return cb
	}

	st := circuitbreaker.Settings{
		Name: name,
		ReadyToTrip: func(counts circuitbreaker.Counts) bool {
			return counts.ConsecutiveFailures > uint32(util.GetIntEnv("CB_MAX_CONSECUTIVE_FAILURES", 5))
		},
		Timeout: time.Duration(util.GetIntEnv("CB_TIMEOUT", 60)) * time.Second,
		OnStateChange: func(name string, from circuitbreaker.State, to circuitbreaker.State) {
			level.Info(s.log).Log(
				util.LogCircuitBreakerEndpoint, name,
				util.LogCircuitBreakerOldStatus, from,
				util.LogCircuitBreakerNewStatus, to,
			)

			if to == circuitbreaker.StateOpen {
				go func() {
					err := s.broker.Publish(context.Background(), name, to.String())
					if err != nil {
						level.Error(s.log).Log(
							util.LogMessage, "failed to publish circuit breaker status",
							util.LogError, err,
							util.LogCircuitBreakerEndpoint, name,
							util.LogCircuitBreakerNewStatus, to.String(),
						)
					}

					requiringEndpoints, err := s.repository.GetMemberOfSet(context.Background(), util.FormRequiringEndpointsKey(name))
					if err != nil {
						level.Error(s.log).Log(
							util.LogMessage, "failed to get requiring endpoints from db",
							util.LogError, err,
							util.LogCircuitBreakerEndpoint, name,
							util.LogCircuitBreakerNewStatus, to.String(),
						)
					}

					for _, ep := range requiringEndpoints {
						endpoint := ep
						go func() {
							err := s.broker.Publish(context.Background(), endpoint, to.String())
							if err != nil {
								level.Error(s.log).Log(
									util.LogMessage, "failed to publish circuit breaker status",
									util.LogError, err,
									util.LogCircuitBreakerEndpoint, endpoint,
									util.LogCircuitBreakerNewStatus, to.String(),
								)
							}
						}()
					}
				}()
			}
		},
	}

	cb := circuitbreaker.NewCircuitBreaker(st)
	s.breakers[name] = cb
	return cb
}
