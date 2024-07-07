package main

import (
	"context"
	"fmt"
	"github.com/daffarg/distributed-cascading-cb/config"
	"github.com/daffarg/distributed-cascading-cb/tracer"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/daffarg/distributed-cascading-cb/broker/kafka"
	"github.com/daffarg/distributed-cascading-cb/endpoint"
	"github.com/daffarg/distributed-cascading-cb/protobuf"
	"github.com/daffarg/distributed-cascading-cb/repository/kvrocks"
	"github.com/daffarg/distributed-cascading-cb/service"
	"github.com/daffarg/distributed-cascading-cb/transport"
	"github.com/daffarg/distributed-cascading-cb/util"
	logkit "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	godotenv.Load()
}

var (
	writer io.Writer
)

func main() {
	logDir := os.Getenv("LOG_DIR")

	if logDir != "" {
		err := os.MkdirAll(os.Getenv("LOG_DIR"), 0755)
		if err != nil {
			fmt.Println("Failed to create log directory:", err)
			os.Exit(1)
		}

		logfile, err := os.OpenFile(filepath.Join(logDir, util.GetEnv("LOG_FILE_NAME", "app.log")), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Printf("error opening log file: %v\n", err)
			os.Exit(1)
		}
		defer logfile.Close()

		writer = logfile
	} else {
		writer = os.Stdout
	}

	var log logkit.Logger
	{
		log = logkit.NewJSONLogger(writer)
		log = logkit.With(log, util.LogTimestamp, logkit.TimestampFormat(time.Now, time.RFC3339), util.LogPath, logkit.DefaultCaller)
	}
	defer level.Info(log).Log(util.LogMessage, "service stopped")

	cbConsumerGroup := os.Getenv("CB_CONSUMER_GROUP")
	if cbConsumerGroup == "" {
		level.Error(log).Log(util.LogError, "please set CB_CONSUMER_GROUP in env variable")
		os.Exit(1)
	}

	var sysLog logkit.Logger
	{
		sysLog = logkit.NewJSONLogger(writer)
		sysLog = logkit.With(sysLog, util.LogTimestamp, logkit.TimestampFormat(time.Now, time.RFC3339), util.LogPath, logkit.Caller(4))
	}

	tracingProvider, err := tracer.NewTracerProvider(
		context.Background(), os.Getenv("SERVICE_NAME"), util.GetEnv("TRACING_BACKEND_URL", "localhost:4317"),
	)
	if err != nil {
		level.Error(log).Log(
			util.LogError, err,
		)
		return
	}

	stopTracingProvider, err := tracingProvider.RegisterAsGlobal()
	if err != nil {
		level.Error(log).Log(
			util.LogError, err,
		)
		return
	}
	defer func() {
		if err = stopTracingProvider(context.TODO()); err != nil {
			level.Error(log).Log(
				util.LogError, err,
			)
			return
		}
	}()

	otelTracer := otel.Tracer(util.GetEnv("SERVICE_NAME", "CB SERVICE"))

	kvRocks, err := kvrocks.NewKVRocksRepository(
		util.GetEnv("KVROCKS_HOST", "127.0.0.1"),
		util.GetEnv("KVROCKS_PORT", "6666"),
		util.GetEnv("KVROCKS_PASSWORD", ""),
		util.GetIntEnv("KVROCKS_DB", 0),
	)
	if err != nil {
		level.Error(log).Log(
			util.LogError, err,
		)
		return
	}

	cbConfig := config.NewConfig()
	err = cbConfig.Read(util.GetEnv("CONFIG_PATH", "config.yaml"))
	if err != nil {
		level.Error(log).Log(
			util.LogError, err,
		)
		return
	}

	kafkaBroker, err := kafka.NewKafkaBroker(
		log,
		util.GetEnv("KAFKA_CONFIG_PATH", "client.properties"),
	)
	if err != nil {
		level.Error(log).Log(
			util.LogError, err,
		)
		return
	}

	circuitBreakerSvc := service.NewCircuitBreakerService(
		log,
		validator.New(),
		kvRocks,
		kafkaBroker,
		&http.Client{
			Timeout:   10 * time.Second,
			Transport: otelhttp.NewTransport(http.DefaultTransport),
		},
		otelTracer,
		cbConfig,
	)

	circuitBreakerEndpoint, err := endpoint.NewCircuitBreakerEndpoint(circuitBreakerSvc, sysLog)
	if err != nil {
		level.Error(log).Log(
			util.LogError, err,
		)
		return
	}

	circuitBreakerServer := transport.NewCircuitBreakerServer(circuitBreakerEndpoint)
	address := fmt.Sprintf("%s:%s", util.GetEnv("SERVICE_IP", "127.0.0.1"), util.GetEnv("SERVICE_PORT", "5320"))

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	lis, errListen := net.Listen("tcp", address)
	if errListen != nil {
		level.Error(log).Log(
			util.LogError, errListen,
		)
		return
	}

	protobuf.RegisterCircuitBreakerServer(grpcServer, circuitBreakerServer)
	reflection.Register(grpcServer)

	// Serve gRPC Server
	level.Info(log).Log(util.LogMessage, fmt.Sprintf("Serving gRPC on %s", address))
	level.Error(log).Log(
		util.LogError,
		grpcServer.Serve(lis),
	)
}
