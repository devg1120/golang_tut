package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerOption func(*zap.Config)

func NewLogger(opts ...LoggerOption) (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	for _, opt := range opts {
		opt(&config)
	}
	return config.Build()
}

type options struct {
	logger *zap.Logger
}

type Option func(*options)

func WithLevel(level zapcore.Level) LoggerOption {
	return func(cfg *zap.Config) {
		cfg.Level = zap.NewAtomicLevelAt(level)
	}
}

func WithOutputPaths(paths []string) LoggerOption {
	return func(cfg *zap.Config) {
		cfg.OutputPaths = paths
	}
}

func WithEncoder(encoder string) LoggerOption {
	return func(cfg *zap.Config) {
		if encoder == "json" {
			cfg.Encoding = "json"
		} else if encoder == "console" {
			cfg.Encoding = "console"
		}
	}
}
