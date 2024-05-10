package service

import (
	"context"
	"github.com/daffarg/distributed-cascading-cb/protobuf"
	"time"

	"github.com/daffarg/distributed-cascading-cb/circuitbreaker"
	"github.com/daffarg/distributed-cascading-cb/util"
	"github.com/go-kit/log/level"
)

func (s *service) getCircuitBreaker(name string) *circuitbreaker.CircuitBreaker {
	if cb, ok := s.breakers[name]; ok {
		return cb
	}

	timeout := time.Duration(util.GetIntEnv("CB_TIMEOUT", 60)) * time.Second
	st := circuitbreaker.Settings{
		Name: name,
		ReadyToTrip: func(counts circuitbreaker.Counts) bool {
			return counts.ConsecutiveFailures >= uint32(util.GetIntEnv("CB_MAX_CONSECUTIVE_FAILURES", 5))
		},
		Timeout: timeout,
		OnStateChange: func(name string, from circuitbreaker.State, to circuitbreaker.State) {
			if to == circuitbreaker.StateOpen {
				go func() {
					level.Info(s.log).Log(
						util.LogCircuitBreakerEndpoint, name,
						util.LogCircuitBreakerOldStatus, from,
						util.LogCircuitBreakerNewStatus, to,
					)

					requiringEndpoints, err := s.repository.GetMemberOfSet(context.Background(), util.FormRequiringEndpointsKey(name))
					if err != nil {
						level.Error(s.log).Log(
							util.LogMessage, "failed to get requiring endpoints from db",
							util.LogError, err,
							util.LogCircuitBreakerEndpoint, name,
							util.LogCircuitBreakerNewStatus, to.String(),
						)
					}

					go func(reqEp []string) {
						for _, ep := range reqEp {
							go func(ep string) {
								encodedTopic := util.EncodeTopic(ep)
								message := &protobuf.Status{
									Endpoint:  ep,
									Status:    to.String(),
									Timeout:   uint32(util.GetIntEnv("CB_TIMEOUT", 60)),
									Timestamp: time.Now().Format(time.RFC3339),
								}
								if err != nil {
									level.Error(s.log).Log(
										util.LogMessage, "failed to marshal circuit breaker status",
										util.LogError, err,
										util.LogCircuitBreakerEndpoint, ep,
										util.LogCircuitBreakerNewStatus, to.String(),
									)
								}

								err = s.broker.Publish(context.Background(), encodedTopic, message)
								if err != nil {
									level.Error(s.log).Log(
										util.LogMessage, "failed to publish circuit breaker status",
										util.LogError, err,
										util.LogCircuitBreakerEndpoint, ep,
										util.LogCircuitBreakerNewStatus, to.String(),
									)
								} else {
									level.Info(s.log).Log(
										util.LogMessage, "published circuit breaker status",
										util.LogStatus, message,
									)
								}
							}(ep)

							err = s.repository.SetWithExp(
								context.Background(),
								util.FormEndpointStatusKey(ep),
								to.String(),
								time.Duration(util.GetIntEnv("CB_TIMEOUT", 60))*time.Second,
							)
							if err != nil {
								level.Error(s.log).Log(
									util.LogMessage, "failed to set circuit breaker status to db",
									util.LogError, err,
									util.LogCircuitBreakerEndpoint, name,
									util.LogCircuitBreakerNewStatus, to.String(),
								)
							}
						}
					}(requiringEndpoints)
				}()
			}
		},
	}

	cb := circuitbreaker.NewCircuitBreaker(st)
	s.breakers[name] = cb
	return cb
}
