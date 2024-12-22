package otel

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func Init(serverName string) {
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)
	tracerProvider, err := newHttpTraceProvider(serverName)
	if err != nil {
		return
	}

	//shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)

	meterProvider, err := newMeterProvider()
	if err != nil {
		return
	}
	otel.SetMeterProvider(meterProvider)
}
