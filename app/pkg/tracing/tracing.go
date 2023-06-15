package tracing

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdk_trace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
	"production_service/pkg/common/core/validator"
	"production_service/pkg/common/errors"
)

type config struct {
	jaegerAgentHost string `validate:"required"`
	jaegerAgentPort string `validate:"required"`
	serviceID       string `validate:"required"`
	serviceName     string `validate:"required"`
	serviceVersion  string `validate:"required"`
	envName         string `validate:"required"`
}

type ConfigParam func(config *config)

func WithJaegerAgentHost(val string) ConfigParam {
	return func(c *config) {
		c.jaegerAgentHost = val
	}
}

func WithJaegerAgentPort(val string) ConfigParam {
	return func(c *config) {
		c.jaegerAgentPort = val
	}
}

func WithServiceID(val string) ConfigParam {
	return func(c *config) {
		c.serviceID = val
	}
}

func WithServiceName(val string) ConfigParam {
	return func(c *config) {
		c.serviceName = val
	}
}

func WithServiceVersion(val string) ConfigParam {
	return func(c *config) {
		c.serviceVersion = val
	}
}

func WithEnvName(val string) ConfigParam {
	return func(c *config) {
		c.envName = val
	}
}

func New(cp ...ConfigParam) (*sdk_trace.TracerProvider, error) {
	cfg := new(config)

	for _, param := range cp {
		param(cfg)
	}

	err := validator.StructValidator(cfg).Validate()
	if err != nil {
		return nil, err
	}

	exporter, err := jaeger.New(
		jaeger.WithAgentEndpoint(
			jaeger.WithAgentHost(cfg.jaegerAgentHost),
			jaeger.WithAgentPort(cfg.jaegerAgentPort),
		),
	)
	if err != nil {
		return nil, errors.Wrap(err, "jaeger.New")
	}

	provider := sdk_trace.NewTracerProvider(
		sdk_trace.WithBatcher(exporter),
		sdk_trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cfg.serviceName),
			semconv.DeploymentEnvironmentKey.String(cfg.envName),
			semconv.ServiceVersionKey.String(cfg.serviceVersion),
			semconv.ServiceInstanceIDKey.String(cfg.serviceID),
		)),
	)

	otel.SetTracerProvider(provider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return provider, nil
}

func SpanEvent(ctx context.Context, name string) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return
	}

	span.AddEvent(name)
}

func Start(ctx context.Context, name string, options ...trace.SpanStartOption) (context.Context, trace.Span) {
	t := otel.Tracer(name)

	ctx, span := t.Start(ctx, name, options...)

	return ctx, span
}

func Continue(ctx context.Context, name string, options ...trace.SpanStartOption) (context.Context, trace.Span) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return ctx, span
	}

	return Start(ctx, name, options...)
}
