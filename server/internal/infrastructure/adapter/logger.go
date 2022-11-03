package adapter

import (
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.SugaredLogger
}

func NewLogger() *Logger {
	l, err := zap.NewProduction(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	return &Logger{
		logger: l.Sugar(),
	}
}

func (l *Logger) EnableDebug() error {
	ld, err := zap.NewDevelopment(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		return err
	}

	l.logger = ld.Sugar()
	return nil
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.logger.Debugw(msg, args...)
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.logger.Warnw(msg, args...)
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.logger.Infow(msg, args...)
}

func (l *Logger) Error(err error, args ...interface{}) {
	l.logger.Errorw(err.Error(), args...)
}

func (l *Logger) Fatal(err error, args ...interface{}) {
	l.logger.Fatalw(err.Error(), args...)
}
