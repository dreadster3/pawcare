package endpoint

import (
	"context"
	"time"

	"github.com/dreadster3/pawcare/services/account/internal/owner/domain"
	"github.com/dreadster3/pawcare/services/account/internal/owner/service"
	"github.com/dreadster3/pawcare/shared/common"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/golang-jwt/jwt/v4"

	"github.com/go-playground/validator/v10"
)

type CreateRequest struct {
	Name        string    `json:"name" validate:"required"`
	DateOfBirth time.Time `json:"date_of_birth" validate:"required"`
}

type CreateResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func makeCreateEndpoint(ownerService service.IOwnerService) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(CreateRequest)
		if !ok {
			return nil, common.ErrCastRequest
		}

		err := validator.New().Struct(req)
		if err != nil {
			return nil, err
		}

		claims, ok := context.Value(kitjwt.JWTClaimsContextKey).(*jwt.StandardClaims)
		if !ok {
			return nil, common.ErrParsingClaims
		}

		userId := domain.UserId(claims.Subject)
		profile := domain.NewOwnerProfile(req.Name, req.DateOfBirth)
		owner, err := ownerService.Create(context, userId, profile)
		if err != nil {
			return nil, err
		}

		return CreateResponse{string(owner.Id), owner.Profile.Name}, nil
	}
}
