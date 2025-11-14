package tracing

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// StartSpan starts a new span
func StartSpan(ctx context.Context, tracer trace.Tracer, spanName string) (context.Context, trace.Span) {
	return tracer.Start(ctx, spanName)
}

// AddEvent adds an event to the span
func AddEvent(span trace.Span, name string, attributes ...attribute.KeyValue) {
	span.AddEvent(name, trace.WithAttributes(attributes...))
}

// SetError marks the span as error
func SetError(span trace.Span, err error) {
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())
}

// SetAttributes sets attributes on the span
func SetAttributes(span trace.Span, attributes ...attribute.KeyValue) {
	span.SetAttributes(attributes...)
}