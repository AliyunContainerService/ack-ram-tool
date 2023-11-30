package log

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

type contextKey struct{}

func Named(name string) *zap.SugaredLogger {
	return Logger.Named(fmt.Sprintf("[%s]", name))
}

func FromContext(ctx context.Context) *zap.SugaredLogger {
	if v, ok := ctx.Value(contextKey{}).(*zap.SugaredLogger); ok {
		return v
	}

	return Logger
}

func NewContext(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, contextKey{}, logger)
}

func IntoContext(ctx context.Context, log *zap.SugaredLogger) context.Context {
	return NewContext(ctx, log)
}
