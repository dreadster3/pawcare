package endpoint

import (
	"context"
	"errors"
	"time"

	"github.com/dreadster3/pawcare/services/account/service"
	"github.com/go-kit/kit/endpoint"

	"github.com/go-playground/validator/v10"
)

type CreateAccountRequest struct {
	Name        string    `json:"name" validate:"required"`
	DateOfBirth time.Time `json:"date_of_birth" validate:"required"`
}

type CreateAccountResponse struct {
	Id string `json:"id"`
}

func makeCreateOwnerEndpoint(svc service.ProfileService) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(CreateAccountRequest)
		if !ok {
			return nil, errors.New("Cannot cast request")
		}

		err := validator.New().Struct(req)
		if err != nil {
			return nil, err
		}

		id, err := svc.CreateAccount(req.Name, req.DateOfBirth)
		if err != nil {
			return nil, err
		}

		return CreateAccountResponse{id}, nil
	}
}
