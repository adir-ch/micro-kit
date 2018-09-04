package service

import (
	"context"
	"log"
)

// CalcService describes the service.
type CalcService interface {
	// Calculate - evaluate a math expression and return the result
	Calculate(ctx context.Context, expr string) (rs float64, err error)
}

type basicCalcService struct{}

func (b *basicCalcService) Calculate(ctx context.Context, expr string) (rs float64, err error) {
	log.Printf("got request to calculate: %s", expr)
	rs, err = eval(expr)
	return rs, err
}

// NewBasicCalcService returns a naive, stateless implementation of CalcService.
func NewBasicCalcService() CalcService {
	return &basicCalcService{}
}

// New returns a CalcService with all of the expected middleware wired in.
func New(middleware []Middleware) CalcService {
	var svc CalcService = NewBasicCalcService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
