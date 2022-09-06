package config

import (
	"fmt"
	"os"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.Logger {
	logLevelEnv := os.Getenv("LOG_LEVEL")
	var logLevel zapcore.Level
	var err error
	if err = logLevel.UnmarshalText([]byte(logLevelEnv)); err != nil {
		logLevel = zapcore.DebugLevel
	}
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, logLevel)
	logger := zap.New(core, zap.AddCaller())
	if err != nil || logLevelEnv == "" {
		logger.Warn("log level is not set or is invalid", zap.Error(err))
	}
	logger.Info(fmt.Sprintf("Logger created, log level: %s", logLevelEnv))
	return logger
}
