package main

import (
	"fmt"
	"github.com/cocoide/golang-design-pattern/singleton_pattern/ctxutils"
	l "github.com/cocoide/golang-design-pattern/singleton_pattern/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"sync"
)

func testSingleton(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	ctx := ctxutils.NewContextWithTraceID(context.Background(), uuid.NewString())
	logger := l.GetLogger(ctx)

	logger.Info("test", zap.Int("logger id", id))
	fmt.Printf("Logger instance %d: %p\n", id, logger)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go testSingleton(&wg, i)
	}

	wg.Wait()
}
