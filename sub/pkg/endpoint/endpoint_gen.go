// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	service "github.com/adir-ch/micro-kit/sub/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	SubEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.SubService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{SubEndpoint: MakeSubEndpoint(s)}
	for _, m := range mdw["Sub"] {
		eps.SubEndpoint = m(eps.SubEndpoint)
	}
	return eps
}
