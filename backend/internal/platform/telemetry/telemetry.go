package telemetry

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

// Setup initializes the global OpenTelemetry tracer provider with an OTLP/HTTP
// trace exporter. It returns a shutdown function that flushes spans on exit.
// Attribute redaction is enforced by instrumentation code; sensitive context is
// never added as a span attribute.
func Setup(ctx context.Context, serviceName, otlpEndpoint string, disabled bool, log *slog.Logger) (func(context.Context) error, error) {
	if disabled {
		log.Info("opentelemetry disabled")
		return func(context.Context) error { return nil }, nil
	}

	exporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpointURL(otlpEndpoint),
		otlptracehttp.WithTimeout(5*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("create otlp trace exporter: %w", err)
	}

	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(serviceName),
		attribute.String("deployment.environment", "development"),
	)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	return tp.Shutdown, nil
}
