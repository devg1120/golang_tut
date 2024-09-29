package main

import (
	l "github.com/cocoide/golang-design-pattern/functional_options_pattern/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	customLogger, err := l.NewLogger(
		l.WithLevel(zapcore.DebugLevel),
		l.WithOutputPaths([]string{"stdout", "logs/app.log"}),
		l.WithEncoder("json"),
	)
	if err != nil {
		panic(err)
	}
	customLogger.Info("custom logger triggered")
}
