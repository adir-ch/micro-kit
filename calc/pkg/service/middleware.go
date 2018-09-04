package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(CalcService) CalcService

type loggingMiddleware struct {
	logger log.Logger
	next   CalcService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a CalcService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next CalcService) CalcService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Calculate(ctx context.Context, expr string) (rs float64, err error) {
	defer func() {
		l.logger.Log("method", "Calculate", "expr", expr, "rs", rs, "err", err)
	}()
	return l.next.Calculate(ctx, expr)
}
