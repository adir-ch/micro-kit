package endpoint

import (
	"context"
	service "github.com/adir-ch/micro-kit/add/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// AddRequest collects the request parameters for the Add method.
type AddRequest struct {
	Numbers []float64 `json:"numbers"`
}

// AddResponse collects the response parameters for the Add method.
type AddResponse struct {
	Rs  float64 `json:"rs"`
	Err error   `json:"err"`
}

// MakeAddEndpoint returns an endpoint that invokes Add on the service.
func MakeAddEndpoint(s service.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		rs, err := s.Add(ctx, req.Numbers)
		return AddResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r AddResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Add implements Service. Primarily useful in a client.
func (e Endpoints) Add(ctx context.Context, numbers []float64) (rs float64, err error) {
	request := AddRequest{Numbers: numbers}
	response, err := e.AddEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddResponse).Rs, response.(AddResponse).Err
}
