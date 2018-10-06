package grpc

import (
	"context"

	endpoint "github.com/adir-ch/micro-kit/sub/pkg/endpoint"
	pb "github.com/adir-ch/micro-kit/sub/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeSubHandler creates the handler logic
func makeSubHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SubEndpoint, decodeSubRequest, encodeSubResponse, options...)
}

// decodeSubResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeSubRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.SubRequest)
	return endpoint.SubRequest{Left: req.Left, Right: req.Right}, nil
}

// encodeSubResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSubResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.SubResponse)
	return &pb.SubReply{Result: res.Rs}, res.Err
}

func (g *grpcServer) Sub(ctx context1.Context, req *pb.SubRequest) (*pb.SubReply, error) {
	_, rep, err := g.sub.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SubReply), nil
}
