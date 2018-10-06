package service

import "context"

// SubService describes the service.
type SubService interface {
	// Add your methods here
	Sub(ctx context.Context, left, right float64) (rs float64, err error)
}

type basicSubService struct{}

func (b *basicSubService) Sub(ctx context.Context, left float64, right float64) (rs float64, err error) {
	rs = left - right
	return rs, nil
}

// NewBasicSubService returns a naive, stateless implementation of SubService.
func NewBasicSubService() SubService {
	return &basicSubService{}
}

// New returns a SubService with all of the expected middleware wired in.
func New(middleware []Middleware) SubService {
	var svc SubService = NewBasicSubService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
