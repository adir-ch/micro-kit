package endpoint

import (
	"context"

	service "github.com/adir-ch/micro-kit/mul/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// MulRequest collects the request parameters for the Mul method.
type MulRequest struct {
	Numbers []float64 `json:"numbers"`
}

// MulResponse collects the response parameters for the Mul method.
type MulResponse struct {
	Rs  float64 `json:"rs"`
	Err error   `json:"err"`
}

// MakeMulEndpoint returns an endpoint that invokes Mul on the service.
func MakeMulEndpoint(s service.MulService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(MulRequest)
		rs, err := s.Mul(ctx, req.Numbers)
		return MulResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r MulResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Mul implements Service. Primarily useful in a client.
func (e Endpoints) Mul(ctx context.Context, numbers []float64) (rs float64, err error) {
	request := MulRequest{
		Numbers: numbers,
	}
	response, err := e.MulEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(MulResponse).Rs, response.(MulResponse).Err
}
