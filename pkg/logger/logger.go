package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func New(loglevel zapcore.Level) *zap.Logger {
	var err error
	//var level zapcore.Level
	//level, err = zapcore.ParseLevel(loglevel)
	//if err != nil {
	//	panic("Parse log level error!!!")
	//}

	logger, err = zap.NewProduction()

	if err != nil {
		panic("Error create logger!!!")
	}
	return logger
}

func CloseLogger() {
	if logger == nil {
		panic("Close nil logger error!!!")
	}
	err := logger.Sync()
	if err != nil {
		return
	}
}

func GetLogger() *zap.Logger {
	if logger == nil {
		panic("logger is nil!!!")
	}
	return logger
}
