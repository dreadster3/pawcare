package transport

import (
	"context"

	"github.com/dreadster3/pawcare/services/account/internal/pet/endpoint"
	"github.com/dreadster3/pawcare/services/account/internal/pet/proto"
	"github.com/dreadster3/pawcare/shared/common"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/transport"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type grpcServer struct {
	proto.UnimplementedPetServiceServer

	getById grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Set, logger log.Logger) proto.PetServiceServer {
	options := []grpc.ServerOption{
		grpc.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		grpc.ServerBefore(kitjwt.GRPCToContext()),
	}
	options = append(options, common.GRPCLoggingServerOptions(logger)...)

	getByIdHandler := grpc.NewServer(
		endpoints.GetByIdEndpoint,
		decodeGetByIdRequest,
		encodeGetByIdResponse,
		options...,
	)

	return &grpcServer{
		getById: getByIdHandler,
	}
}

func (srv *grpcServer) GetById(ctx context.Context, req *proto.GetByIdRequest) (*proto.GetPetResponse, error) {
	_, res, err := srv.getById.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.(*proto.GetPetResponse), nil
}

func decodeGetByIdRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(*proto.GetByIdRequest)
	if !ok {
		return nil, common.ErrCastRequest
	}

	return endpoint.GetByIdRequest{
		Id: req.Id,
	}, nil
}

func encodeGetByIdResponse(_ context.Context, res interface{}) (interface{}, error) {
	response, ok := res.(endpoint.GetResponse)
	if !ok {
		return nil, common.ErrCastResponse
	}

	return &proto.GetPetResponse{
		Id:          response.Id,
		Name:        response.Name,
		Species:     response.Species,
		DateOfBirth: timestamppb.New(response.DateOfBirth),
		Weight:      response.Weight,
		Breed:       response.Breed,
		Gender:      response.Gender,
		Err:         common.Err2Str(response.Err),
	}, nil
}
