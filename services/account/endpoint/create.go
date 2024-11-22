package endpoint

import (
	"context"
	"time"

	"github.com/dreadster3/pawcare/services/account/service"
	"github.com/go-kit/kit/endpoint"
)

type CreateAccountRequest struct {
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

type createAccountResponse struct {
	Id string `json:"id"`
}

func makeCreateOwnerEndpoint(svc service.ProfileService) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAccountRequest)
		id, err := svc.CreateAccount(req.Name, req.DateOfBirth)
		if err != nil {
			return nil, err
		}

		return createAccountResponse{id}, nil
	}
}
