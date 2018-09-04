package service

import "context"

// AddService describes the service.
type AddService interface {
	// Add - returns the sum of numbers
	Add(ctx context.Context, numbers []float64) (rs float64, err error)
}

type basicAddService struct{}

func (b *basicAddService) Add(ctx context.Context, numbers []float64) (rs float64, err error) {
	for _, n := range numbers {
		rs += n
	}
	return rs, err
}

// NewBasicAddService returns a naive, stateless implementation of AddService.
func NewBasicAddService() AddService {
	return &basicAddService{}
}

// New returns a AddService with all of the expected middleware wired in.
func New(middleware []Middleware) AddService {
	var svc AddService = NewBasicAddService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
