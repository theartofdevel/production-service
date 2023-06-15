package logging

import (
	"context"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func AnyField(key string, val interface{}) zap.Field {
	return zap.Any(key, val)
}

func ErrorField(err error) zap.Field {
	return zap.Error(err)
}

func StringField(key, val string) zap.Field {
	return zap.String(key, val)
}

func IntField(key string, val int) zap.Field {
	return zap.Int(key, val)
}

func DurationField(key string, val time.Duration) zap.Field {
	return zap.Duration(key, val)
}

func Int64Field(key string, val int64) zap.Field {
	return zap.Int64(key, val)
}

func StringsField(key string, val []string) zap.Field {
	return zap.Strings(key, val)
}

func BoolField(key string, val bool) zap.Field {
	return zap.Bool(key, val)
}

var (
	defaultLogger = NewLogger()
	atomicLevel   = zap.NewAtomicLevel()
)

func NewLogger() *zap.Logger {
	var encCfg zapcore.EncoderConfig

	var encoder zapcore.Encoder

	encCfg = zap.NewProductionEncoderConfig()
	encCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encCfg.LevelKey = zapcore.DebugLevel.String()

	encCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encCfg.TimeKey = "timestamp"

	encoder = zapcore.NewConsoleEncoder(encCfg)

	l := zap.New(
		zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), atomicLevel),
	)
	l.WithOptions(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	return l
}

func GetLogger() *zap.Logger {
	return defaultLogger
}

func L(ctx context.Context) *zap.Logger {
	return LoggerFromContext(ctx)
}

func SetLevel(level string) {
	var lvl zapcore.Level

	switch strings.ToLower(level) {
	case "info":
		lvl = zapcore.InfoLevel
	case "error":
		lvl = zapcore.ErrorLevel
	default:
		lvl = zapcore.DebugLevel
	}

	atomicLevel.SetLevel(lvl)
}

func WithField(ctx context.Context, field zap.Field) *zap.Logger {
	return LoggerFromContext(ctx).With(field)
}

func WithFields(ctx context.Context, fields ...zap.Field) *zap.Logger {
	return LoggerFromContext(ctx).With(fields...)
}

func WithError(ctx context.Context, err error) *zap.Logger {
	return LoggerFromContext(ctx).With(zap.Error(err))
}
