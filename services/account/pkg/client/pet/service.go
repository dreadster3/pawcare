package pet

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type PetService interface {
	GetById(ctx context.Context, id PetId) (Pet, error)
}

func NewPetService(baseUrl string) (PetService, error) {
	client, err := NewHTTPClient(baseUrl)
	if err != nil {
		return nil, err
	}

	return &petClient{
		getById: client.GetById,
	}, nil
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
		return Pet{}, nil
	}

	r := resp.(GetByIdResponse)
	if r.Err != "" {
		return Pet{}, errors.New(r.Err)
	}

	return r.Profile, nil
}
