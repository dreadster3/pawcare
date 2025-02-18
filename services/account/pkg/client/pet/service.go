package pet

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"google.golang.org/grpc"
)

type IPetService interface {
	GetById(ctx context.Context, id PetId) (Pet, error)
}

func NewPetService(conn *grpc.ClientConn, logger log.Logger) IPetService {
	client := newGRPCClient(conn)

	var svc IPetService
	svc = &petClient{
		getById: client.GetById,
	}
	svc = newLoggingMiddleware(logger)(svc)
	return svc
}

type petClient struct {
	getById endpoint.Endpoint
}

func (c *petClient) GetById(ctx context.Context, id PetId) (Pet, error) {
	req := GetByIdRequest{
		Id: string(id),
	}
	resp, err := c.getById(ctx, req)
	if err != nil {
		return Pet{}, err
	}

	r := resp.(GetByIdResponse)
	if r.Err != "" {
		return Pet{}, errors.New(r.Err)
	}

	return r.Profile, nil
}
