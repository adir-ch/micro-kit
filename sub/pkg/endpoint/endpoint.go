package endpoint

import (
	"context"
	service "github.com/adir-ch/micro-kit/sub/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// SubRequest collects the request parameters for the Sub method.
type SubRequest struct {
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
}

// SubResponse collects the response parameters for the Sub method.
type SubResponse struct {
	Rs  float64 `json:"rs"`
	Err error   `json:"err"`
}

// MakeSubEndpoint returns an endpoint that invokes Sub on the service.
func MakeSubEndpoint(s service.SubService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SubRequest)
		rs, err := s.Sub(ctx, req.Left, req.Right)
		return SubResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r SubResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Sub implements Service. Primarily useful in a client.
func (e Endpoints) Sub(ctx context.Context, left float64, right float64) (rs float64, err error) {
	request := SubRequest{
		Left:  left,
		Right: right,
	}
	response, err := e.SubEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SubResponse).Rs, response.(SubResponse).Err
}
