package log

import "go.uber.org/zap"

type ProviderLogWrapper struct {
	ZP *zap.SugaredLogger
}

func (l *ProviderLogWrapper) Info(msg string) {
	l.ZP.Info(msg)
}

func (l *ProviderLogWrapper) Debug(msg string) {
	l.ZP.Debug(msg)
}

func (l *ProviderLogWrapper) Error(err error, msg string) {
	l.ZP.Error(msg)
}
