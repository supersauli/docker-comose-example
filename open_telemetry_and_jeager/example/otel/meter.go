package otel

import (
	"time"

	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/sdk/metric"
)

func newMeterProvider() (*metric.MeterProvider, error) {

	metricExporter, err := stdoutmetric.New()
	if err != nil {
		return nil, err
	}

	// 创建一个 MeterProvider，配置处理器、导出器等

	mp := metric.NewMeterProvider(
		//metric.WithReader(metric.NewPeriodicReader(exp, metric.WithInterval(3*time.Second))),
		metric.WithReader(metric.NewPeriodicReader(metricExporter, metric.WithInterval(3*time.Second))),

		//metric.WithResource(resource.NewWithAttributes(
		//	semconv.SchemaURL,
		//	semconv.ServiceNameKey.String("your_service_name"),
		//	semconv.ServiceVersionKey.String("1.0.0"),
		//)
	)
	//metric.WithBatcher(
	//	// 创建一个 OTLP exporter 并配置其发送数据的地址等
	//	otlp.NewExporter(
	//		otlpmetric.WithInsecure(),
	//		otlpmetric.WithEndpoint("localhost:4317"),
	//	),
	//),
	//)

	return mp, nil
}
