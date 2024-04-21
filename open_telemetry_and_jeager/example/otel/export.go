package otel

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/trace"
)

func newHttpOTLPExporter(ctx context.Context) (trace.SpanExporter, error) {
	return otlptracehttp.New(ctx, otlptracehttp.WithEndpointURL("http://192.168.68.166:24318"))
}
