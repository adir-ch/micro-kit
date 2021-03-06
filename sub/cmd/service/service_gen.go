// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint "github.com/adir-ch/micro-kit/sub/pkg/endpoint"
	service "github.com/adir-ch/micro-kit/sub/pkg/service"
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	grpc "github.com/go-kit/kit/transport/grpc"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initGRPCHandler(endpoints, g)
	return g
}
func defaultGRPCOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]grpc.ServerOption {
	options := map[string][]grpc.ServerOption{"Sub": {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "Sub", logger))}}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["Sub"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Sub")), endpoint.InstrumentingMiddleware(duration.With("method", "Sub"))}
}
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"Sub"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
