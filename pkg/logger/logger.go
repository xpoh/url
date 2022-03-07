// Copyright 2022. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package logger implement logger function
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

// New is a constructor
func New(loglevel zapcore.Level) *zap.Logger {
	var err error

	logger, err = zap.NewProduction()

	if err != nil {
		panic("Error create logger!!!")
	}
	return logger
}

// CloseLogger use for gracefull close logger
func CloseLogger() {
	if logger == nil {
		panic("Close nil logger error!!!")
	}
	err := logger.Sync()
	if err != nil {
		return
	}
}

// GetLogger return a global logger object
func GetLogger() *zap.Logger {
	if logger == nil {
		panic("logger is nil!!!")
	}
	return logger
}
