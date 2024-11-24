package transport

import (
	"context"
	"time"

	"github.com/dreadster3/pawcare/services/account/endpoint"
	"github.com/dreadster3/pawcare/services/account/proto"
	"github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type grpcServer struct {
	proto.UnimplementedAccountServiceServer

	create grpctransport.Handler
}

func NewGRPCServer(endpoints endpoint.Set, logger log.Logger) proto.AccountServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	return &grpcServer{
		create: grpctransport.NewServer(endpoints.CreateAccountEndpoint, decodeCreateRequest, encodeCreateResponse, options...),
	}
}

func (srv *grpcServer) Create(ctx context.Context, req *proto.AccountServiceCreateRequest) (*proto.AccountServiceCreateResponse, error) {
	_, res, err := srv.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*proto.AccountServiceCreateResponse), nil
}

func decodeCreateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.AccountServiceCreateRequest)
	return endpoint.CreateAccountRequest{Name: req.Name, DateOfBirth: time.Now()}, nil
}

func encodeCreateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.CreateAccountResponse)
	return &proto.AccountServiceCreateResponse{Id: resp.Id}, nil
}
