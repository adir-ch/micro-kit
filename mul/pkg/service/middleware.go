package service

import (
	"context"
	"fmt"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(MulService) MulService

type loggingMiddleware struct {
	logger log.Logger
	next   MulService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a MulService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next MulService) MulService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Mul(ctx context.Context, numbers []float64) (rs float64, err error) {
	defer func() {
		l.logger.Log("method", "Mul", "numbers", fmt.Sprintf("%v", numbers), "rs", rs, "err", err)
	}()
	return l.next.Mul(ctx, numbers)
}
