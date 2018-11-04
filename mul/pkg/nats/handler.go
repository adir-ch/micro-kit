package nats

import (
	"context"
	"encoding/json"

	endpoint "github.com/adir-ch/micro-kit/mul/pkg/endpoint"
	natstransport "github.com/go-kit/kit/transport/nats"
	nats "github.com/nats-io/go-nats"
)

//  NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewNATSHandler(endpoints endpoint.Endpoints, options []natstransport.SubscriberOption) *natstransport.Subscriber {
	s := natstransport.NewSubscriber(endpoints.MulEndpoint, decodeMulRequest, natstransport.EncodeJSONResponse, options...)
	return s
}

// decodeMulResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the NATS message request body.
func decodeMulRequest(ctx context.Context, msg *nats.Msg) (interface{}, error) {
	req := endpoint.MulRequest{}
	err := json.Unmarshal(msg.Data, &req)
	return req, err
}
