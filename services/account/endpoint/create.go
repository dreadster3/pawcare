package endpoint

import (
	"context"
	"errors"
	"time"

	"github.com/dreadster3/pawcare/services/account/service"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/golang-jwt/jwt/v4"

	"github.com/go-playground/validator/v10"
)

type CreateAccountRequest struct {
	Name        string    `json:"name" validate:"required"`
	DateOfBirth time.Time `json:"date_of_birth" validate:"required"`
}

type CreateAccountResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func makeCreateOwnerEndpoint(accountService service.IAccountService) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(CreateAccountRequest)
		if !ok {
			return nil, errors.New("Cannot cast request")
		}

		err := validator.New().Struct(req)
		if err != nil {
			return nil, err
		}

		claims, ok := context.Value(kitjwt.JWTClaimsContextKey).(*jwt.StandardClaims)
		if !ok {
			return nil, errors.New("Error parsing claims")
		}

		account, err := accountService.CreateAccount(context, claims.Subject, req.Name, req.DateOfBirth)
		if err != nil {
			return nil, err
		}

		return CreateAccountResponse{string(account.Owner.Id), account.Owner.Name}, nil
	}
}
