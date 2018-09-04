package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(AddService) AddService

type loggingMiddleware struct {
	logger log.Logger
	next   AddService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AddService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AddService) AddService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Add(ctx context.Context, numbers []float64) (rs float64, err error) {
	defer func() {
		l.logger.Log("method", "Add", "numbers", numbers, "rs", rs, "err", err)
	}()
	return l.next.Add(ctx, numbers)
}
