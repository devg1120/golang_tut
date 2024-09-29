package logger

import (
	"github.com/cocoide/golang-design-pattern/singleton_pattern/ctxutils"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"sync"
)

var (
	logger *zap.Logger
	once   sync.Once
)

// Generate singleton logger instance.
func GetLogger(ctx context.Context) *zap.Logger {
	once.Do(func() {
		var err error
		config := zap.NewProductionConfig()
		config.Encoding = "json"
		config.OutputPaths = []string{"stdout", "logs/app.log"}
		logger, err = config.Build()
		if err != nil {
			panic(err)
		}
	})
	if traceID, ok := ctxutils.GetTraceIDFromContext(ctx); ok {
		return logger.With(zap.String("trace_id", traceID))
	}
	return logger
}
