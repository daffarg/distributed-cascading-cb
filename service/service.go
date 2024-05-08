package service

import (
	"context"
	"github.com/daffarg/distributed-cascading-cb/broker"
	"github.com/daffarg/distributed-cascading-cb/circuitbreaker"
	"github.com/daffarg/distributed-cascading-cb/config"
	"github.com/daffarg/distributed-cascading-cb/repository"
	"github.com/daffarg/distributed-cascading-cb/util"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/go-playground/validator/v10"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

type CircuitBreakerService interface {
	General(ctx context.Context, req *GeneralRequest) (*Response, error)
	Get(ctx context.Context, req *GetRequest) (*Response, error)
	Post(ctx context.Context, req *PostRequest) (*Response, error)
	Put(ctx context.Context, req *PutRequest) (*Response, error)
	Delete(ctx context.Context, req *DeleteRequest) (*Response, error)
}

type service struct {
	log          log.Logger
	validator    *validator.Validate
	repository   repository.Repository
	broker       broker.MessageBroker
	breakers     map[string]*circuitbreaker.CircuitBreaker
	httpClient   *http.Client
	tracer       trace.Tracer
	config       *config.Config
	subscribeMap map[string]bool
}

func NewCircuitBreakerService(
	log log.Logger,
	validator *validator.Validate,
	repository repository.Repository,
	broker broker.MessageBroker,
	httpClient *http.Client,
	tracer trace.Tracer,
	config *config.Config,
) CircuitBreakerService {
	svc := &service{
		log:          log,
		validator:    validator,
		repository:   repository,
		broker:       broker,
		breakers:     make(map[string]*circuitbreaker.CircuitBreaker),
		httpClient:   httpClient,
		tracer:       tracer,
		config:       config,
		subscribeMap: make(map[string]bool),
	}

	svc.initConfig(context.Background())
	err := svc.initSubscribe(context.Background())
	if err != nil {
		level.Error(svc.log).Log(
			util.LogMessage, "failed to init subscribe",
			util.LogError, err,
		)
	}

	return svc
}
