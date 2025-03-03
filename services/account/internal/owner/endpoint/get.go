package endpoint

import (
	"context"
	"time"

	"github.com/dreadster3/pawcare/services/account/internal/owner/service"
	"github.com/go-kit/kit/endpoint"
)

type GetResponse struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Err         error     `json:"-"`
}

func (r GetResponse) Failed() error {
	return r.Err
}

func makeGetEndpoint(ownerService service.IOwnerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		owner, err := ownerService.Get(ctx)
		if err != nil {
			return GetResponse{Err: err}, nil
		}

		return GetResponse{
			Id:          string(owner.Id),
			Name:        owner.Profile.Name,
			DateOfBirth: owner.Profile.DateOfBirth,
		}, nil
	}
}
