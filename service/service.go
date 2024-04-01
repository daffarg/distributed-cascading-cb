package service

import (
	"context"

	"github.com/daffarg/distributed-cascading-cb/broker"
	"github.com/daffarg/distributed-cascading-cb/circuitbreaker"
	"github.com/daffarg/distributed-cascading-cb/repository"
	"github.com/go-kit/log"
	"github.com/go-playground/validator/v10"
)

type CircuitBreakerService interface {
	GeneralRequest(ctx context.Context, req GeneralRequestReq) (Response, error)
}

type service struct {
	log                      log.Logger
	validator                *validator.Validate
	repository               repository.Repository
	broker                   broker.MessageBroker
	breakers                 map[string]*circuitbreaker.CircuitBreaker
	isRequiringEndpointAdded map[string]bool // true if the requiring endpoint present in the set
}

func NewCircuitBreakerService(log log.Logger, validator *validator.Validate, repository repository.Repository, broker broker.MessageBroker) CircuitBreakerService {
	return &service{
		log:                      log,
		validator:                validator,
		repository:               repository,
		broker:                   broker,
		breakers:                 make(map[string]*circuitbreaker.CircuitBreaker),
		isRequiringEndpointAdded: make(map[string]bool),
	}
}
