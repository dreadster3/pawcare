package endpoint

import (
	"context"
	"time"

	"github.com/dreadster3/pawcare/services/account/internal/owner/domain"
	"github.com/dreadster3/pawcare/services/account/internal/owner/service"
	"github.com/dreadster3/pawcare/shared/common"
	"github.com/go-kit/kit/endpoint"

	"github.com/go-playground/validator/v10"
)

type CreateRequest struct {
	Name        string    `json:"name" validate:"required"`
	DateOfBirth time.Time `json:"date_of_birth" validate:"required"`
}

type CreateResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Err  error  `json:"-"`
}

func (r CreateResponse) Failed() error {
	return r.Err
}

func makeCreateEndpoint(ownerService service.IOwnerService) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(CreateRequest)
		if !ok {
			return CreateResponse{Err: common.ErrCastRequest}, nil
		}

		err := validator.New().Struct(req)
		if err != nil {
			return CreateResponse{Err: err}, nil
		}

		profile := domain.NewOwnerProfile(req.Name, req.DateOfBirth)
		owner, err := ownerService.Create(context, profile)
		if err != nil {
			return CreateResponse{Err: err}, nil
		}

		return CreateResponse{
			Id:   string(owner.Id),
			Name: owner.Profile.Name,
		}, nil
	}
}
