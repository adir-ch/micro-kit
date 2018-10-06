package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(SubService) SubService

type loggingMiddleware struct {
	logger log.Logger
	next   SubService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a SubService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next SubService) SubService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Sub(ctx context.Context, left float64, right float64) (rs float64, err error) {
	defer func() {
		l.logger.Log("method", "Sub", "left", left, "right", right, "rs", rs, "err", err)
	}()
	return l.next.Sub(ctx, left, right)
}
