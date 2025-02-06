package ownerendpoints

import (
	"context"
	"time"

	"github.com/dreadster3/pawcare/services/account/endpoints"
	ownerservice "github.com/dreadster3/pawcare/services/account/service/owner"
	"github.com/dreadster3/pawcare/services/account/valueobjects"
	"github.com/dreadster3/pawcare/services/auth"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/golang-jwt/jwt/v4"

	"github.com/go-playground/validator/v10"
)

type CreateOwnerRequest struct {
	Name        string    `json:"name" validate:"required"`
	DateOfBirth time.Time `json:"date_of_birth" validate:"required"`
}

type CreateOwnerResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func MakeCreateOwnerEndpoint(ownerService ownerservice.IOwnerService) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(CreateOwnerRequest)
		if !ok {
			return nil, endpoints.ErrCastRequest
		}

		err := validator.New().Struct(req)
		if err != nil {
			return nil, err
		}

		claims, ok := context.Value(kitjwt.JWTClaimsContextKey).(*jwt.StandardClaims)
		if !ok {
			return nil, endpoints.ErrParsingClaims
		}

		userId := auth.UserId(claims.Subject)
		profile := valueobjects.NewOwnerProfile(req.Name, req.DateOfBirth)
		owner, err := ownerService.Create(context, userId, profile)
		if err != nil {
			return nil, err
		}

		return CreateOwnerResponse{string(owner.Id), owner.Profile.Name}, nil
	}
}
