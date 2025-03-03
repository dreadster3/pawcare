package transport

import (
	"context"

	"github.com/dreadster3/pawcare/services/account/internal/owner/endpoint"
	"github.com/dreadster3/pawcare/services/account/internal/owner/proto"
	"github.com/dreadster3/pawcare/shared/common"
	"github.com/dreadster3/pawcare/shared/utils"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/transport"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type grpcServer struct {
	proto.UnimplementedOwnerServiceServer

	get grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Set, logger log.Logger) proto.OwnerServiceServer {
	options := []grpc.ServerOption{
		grpc.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		grpc.ServerBefore(kitjwt.GRPCToContext()),
		grpc.ServerBefore(utils.RequestIdGRPCToContext()),
	}
	options = append(options, common.GRPCLoggingServerOptions(logger)...)

	getHandler := grpc.NewServer(
		endpoints.GetEndpoint,
		common.GRPCDecodeNoBody,
		encodeGetResponse,
		options...,
	)

	return &grpcServer{
		get: getHandler,
	}
}

func (srv *grpcServer) Get(ctx context.Context, req *emptypb.Empty) (*proto.GetOwnerResponse, error) {
	_, res, err := srv.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.(*proto.GetOwnerResponse), nil
}

func encodeGetResponse(_ context.Context, res interface{}) (interface{}, error) {
	response, ok := res.(endpoint.GetResponse)
	if !ok {
		return nil, common.ErrCastResponse
	}

	return &proto.GetOwnerResponse{
		Id:          response.Id,
		Name:        response.Name,
		DateOfBirth: response.DateOfBirth.String(),
		Err:         common.Err2Str(response.Err),
	}, nil
}
