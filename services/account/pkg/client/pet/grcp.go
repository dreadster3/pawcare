package pet

import (
	"context"
	"errors"
	"time"

	"github.com/dreadster3/pawcare/services/account/internal/pet/proto"
	"github.com/dreadster3/pawcare/shared/proxy"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

type Set struct {
	GetById endpoint.Endpoint
}

func newGRPCClient(conn *grpc.ClientConn) Set {
	options := []kitgrpc.ClientOption{
		kitgrpc.ClientBefore(kitjwt.ContextToGRPC()),
	}

	var getById endpoint.Endpoint
	{
		getById = kitgrpc.NewClient(conn, "proto.PetService", "GetById", encodeGetByIdRequest, decodeGetByIdResponse, proto.GetPetResponse{}, options...).Endpoint()
		getById = proxy.Retry(3, 250*time.Millisecond)(getById)
	}

	return Set{
		GetById: getById,
	}
}

func encodeGetByIdRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(GetByIdRequest)
	return &proto.GetByIdRequest{
		Id: req.Id,
	}, nil
}

func decodeGetByIdResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp := reply.(*proto.GetPetResponse)

	if resp.Err != "" {
		return nil, errors.New(resp.Err)
	}

	return GetByIdResponse{
		Profile: Pet{
			Id:          PetId(resp.Id),
			OwnerId:     OwnerId(resp.OwnerId),
			Name:        resp.Name,
			Breed:       resp.Breed,
			Species:     resp.Species,
			DateOfBirth: resp.DateOfBirth.AsTime(),
			Weight:      resp.Weight,
			Gender:      resp.Gender,
		},
	}, nil
}

type GetByIdRequest struct {
	Id string `json:"id"`
}

type GetByIdResponse struct {
	Profile Pet
	Err     string
}
