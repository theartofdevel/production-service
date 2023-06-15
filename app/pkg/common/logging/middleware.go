package logging

import (
	"net/http"

	"go.opentelemetry.io/otel/trace"
	"production_service/pkg/tracing"
)

const (
	requestIDLogKey = "request_id"
	traceIDLogKey   = "trace_id"
	spanIDLogKey    = "span_id"
)

func Middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		newLogger := NewLogger()
		newLogger = newLogger.With(StringField("endpoint", r.URL.RequestURI()))

		if span := trace.SpanContextFromContext(ctx); span.HasTraceID() {
			newLogger = newLogger.With(StringField(traceIDLogKey, span.TraceID().String()))
			tracing.TraceVal(ctx, traceIDLogKey, span.TraceID().String())
			newLogger = newLogger.With(StringField(spanIDLogKey, span.TraceID().String()))
		}

		ctx = ContextWithLogger(ctx, newLogger)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
