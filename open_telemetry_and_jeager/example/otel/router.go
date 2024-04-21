package otel

import (
	"go.opentelemetry.io/otel/trace"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func Add(c *gin.Context) {
	tracer := otel.Tracer("Add")
	ctx, span := tracer.Start(c.Request.Context(), "Add")
	_ = ctx
	defer span.End()
	span.SetAttributes(attribute.String("key", "value"))
	span.AddEvent("res", trace.WithAttributes(attribute.Int("AddEvent roll", 1)))
	c.JSON(200, gin.H{"state": "ok"})
}
