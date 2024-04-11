package tracer

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"time"
)

type TracerProvider struct {
	serviceName string
	exporterURL string
	provider    *trace.TracerProvider
}

func NewTracerProvider(ctx context.Context, serviceName, exporterURL string) (*TracerProvider, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	e, err := otlptrace.New(ctx, otlptracegrpc.NewClient(
		otlptracegrpc.WithEndpoint(exporterURL),
		otlptracegrpc.WithInsecure(),
	))
	if err != nil {
		return nil, err
	}

	r := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(serviceName),
		semconv.ServiceVersionKey.String("0.0.1"),
	)

	s := trace.AlwaysSample()

	tracerProvider := trace.NewTracerProvider(
		trace.WithSampler(s),
		trace.WithBatcher(e),
		trace.WithResource(r),
	)

	return &TracerProvider{
		serviceName: serviceName,
		exporterURL: exporterURL,
		provider:    tracerProvider,
	}, nil
}

func (p *TracerProvider) RegisterAsGlobal() (func(ctx context.Context) error, error) {
	otel.SetTracerProvider(p.provider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return p.provider.Shutdown, nil
}
