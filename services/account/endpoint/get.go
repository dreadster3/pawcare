package endpoint

import (
	"context"
	"errors"

	"github.com/dreadster3/pawcare/services/account/service"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/golang-jwt/jwt/v4"

	"github.com/go-playground/validator/v10"
)

type GetAccountRequest struct {
}

type GetAccountResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func makeGetAccountEndpoint(accountService service.IAccountService) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(GetAccountRequest)
		if !ok {
			return nil, errors.New("cannot cast request")
		}

		err := validator.New().Struct(req)
		if err != nil {
			return nil, err
		}

		claims, ok := context.Value(kitjwt.JWTClaimsContextKey).(*jwt.StandardClaims)
		if !ok {
			return nil, errors.New("error parsing claims")
		}

		account, err := accountService.GetAccount(context, claims.Subject)
		if err != nil {
			return nil, err
		}

		return CreateAccountResponse{string(account.Owner.Id), account.Owner.Name}, nil
	}
}
