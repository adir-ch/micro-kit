package service

import "context"

// MulService describes the service.
type MulService interface {
	Mul(ctx context.Context, numbers []float64) (rs float64, err error)
}

type basicMulService struct{}

func (b *basicMulService) Mul(ctx context.Context, numbers []float64) (rs float64, err error) {
	for ind, num := range numbers {
		if ind == 0 {
			rs = num
		} else {
			rs = rs * num
		}

	}
	return rs, err
}

// NewBasicMulService returns a naive, stateless implementation of MulService.
func NewBasicMulService() MulService {
	return &basicMulService{}
}

// New returns a MulService with all of the expected middleware wired in.
func New(middleware []Middleware) MulService {
	var svc MulService = NewBasicMulService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
