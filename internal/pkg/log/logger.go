package log

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type loggerKeyType int

const loggerKey loggerKeyType = iota

var (
	atomic zap.AtomicLevel
	logger *zap.SugaredLogger
)

func init() {
	atomic = zap.NewAtomicLevel()

	zl := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.Lock(os.Stdout),
		atomic,
	))

	logger = zl.Sugar()
}

func NewContext(ctx context.Context, fields ...interface{}) context.Context {
	return context.WithValue(ctx, loggerKey, WithContext(ctx).With(fields...))
}

func WithContext(ctx context.Context) *zap.SugaredLogger {
	if ctxLogger, ok := ctx.Value(loggerKey).(*zap.SugaredLogger); ok {
		return ctxLogger
	}

	return logger
}
