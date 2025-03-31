package logger

import (
	"go.uber.org/zap"
)

type Logger interface {
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Infow(msg string, args ...interface{})
}

var Log Logger

func GetLogger() Logger {
	if Log == nil {
		zapLogger, _ := zap.NewProduction()
		defer zapLogger.Sync()
		sugar := zapLogger.Sugar()
		Log = sugar
	}
	return Log
}
