package endpoint

import (
	"context"
	service "github.com/adir-ch/micro-kit/calc/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// CalculateRequest collects the request parameters for the Calculate method.
type CalculateRequest struct {
	Expr string `json:"expr"`
}

// CalculateResponse collects the response parameters for the Calculate method.
type CalculateResponse struct {
	Rs  float64 `json:"rs"`
	Err error   `json:"err"`
}

// MakeCalculateEndpoint returns an endpoint that invokes Calculate on the service.
func MakeCalculateEndpoint(s service.CalcService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CalculateRequest)
		rs, err := s.Calculate(ctx, req.Expr)
		return CalculateResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r CalculateResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Calculate implements Service. Primarily useful in a client.
func (e Endpoints) Calculate(ctx context.Context, expr string) (rs float64, err error) {
	request := CalculateRequest{Expr: expr}
	response, err := e.CalculateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CalculateResponse).Rs, response.(CalculateResponse).Err
}
