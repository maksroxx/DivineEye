package logger

import (
	"go.uber.org/zap"
)

type Logger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Sync() error
}

type ZapLogger struct {
	Logger *zap.Logger
}

func NewZapLogger() *ZapLogger {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return &ZapLogger{Logger: l}
}

func (l *ZapLogger) Info(msg string, fields ...zap.Field) {
	l.Logger.Info(msg, fields...)
}

func (l *ZapLogger) Error(msg string, fields ...zap.Field) {
	l.Logger.Error(msg, fields...)
}

func (l *ZapLogger) Sync() error {
	return l.Logger.Sync()
}
