package main

import (
	"fmt"
	"io"
	"net"
	"os"
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
	writer = os.Stdout

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

	circuitBreakerSvc := service.NewCircuitBreakerService(log, validator.New(), kvrocks.NewKVRocksRepository(
		util.GetEnv("KVROCKS_HOST", "127.0.0.1"),
		util.GetEnv("KVROCKS_PORT", "6666"),
		util.GetEnv("KVROCKS_PASSWORD", ""),
		util.GetIntEnv("KVROCKS_DB", 0),
	),
		kafka.NewKafkaBroker(
			log,
			util.GetEnv("KAFKA_ADDRESS", "127.0.0.1:9092"),
		),
	)

	var sysLog logkit.Logger
	{
		sysLog = logkit.NewJSONLogger(writer)
		sysLog = logkit.With(sysLog, util.LogTimestamp, logkit.TimestampFormat(time.Now, time.RFC3339), util.LogPath, logkit.Caller(4))
	}

	circuitBreakerEndpoint, err := endpoint.NewCircuitBreakerEndpoint(circuitBreakerSvc, sysLog)
	if err != nil {
		level.Error(log).Log(
			util.LogError, err,
		)
		return
	}

	circuitBreakerServer := transport.NewCircuitBreakerServer(circuitBreakerEndpoint)
	address := fmt.Sprintf("%s:%s", util.GetEnv("SERVICE_IP", "127.0.0.1"), util.GetEnv("SERVICE_PORT", "5320"))

	grpcServer := grpc.NewServer()

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
