package logging

import (
	"context"

	"go.uber.org/zap"
)

// https://dev.to/ilyakaznacheev/where-to-place-logger-in-golang-13o3

type ctxLogger struct{}

// ContextWithLogger adds logger to context
func ContextWithLogger(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, l)
}

// LoggerFromContext returns logger from context
func LoggerFromContext(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(ctxLogger{}).(*zap.Logger); ok {
		return l
	}

	return NewLogger()
}
