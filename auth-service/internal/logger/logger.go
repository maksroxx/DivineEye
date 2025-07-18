package logger

import "go.uber.org/zap"

type Logger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Sync() error
}

type ZapLogger struct {
	zap *zap.Logger
}

func NewZapLogger() *ZapLogger {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return &ZapLogger{zap: l}
}

func (l *ZapLogger) Info(msg string, fields ...zap.Field) {
	l.zap.Info(msg, fields...)
}

func (l *ZapLogger) Error(msg string, fields ...zap.Field) {
	l.zap.Error(msg, fields...)
}

func (l *ZapLogger) Sync() error {
	return l.zap.Sync()
}
