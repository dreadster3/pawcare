package endpoint

import (
	"context"
	"time"

	ownerdomain "github.com/dreadster3/pawcare/services/account/internal/owner/domain"
	"github.com/dreadster3/pawcare/services/account/internal/pet/domain"
	petservice "github.com/dreadster3/pawcare/services/account/internal/pet/service"
	"github.com/dreadster3/pawcare/shared/common"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type CreateRequest struct {
	Name        string         `json:"name" validate:"required"`
	DateOfBirth time.Time      `json:"date_of_birth" validate:"required"`
	Species     string         `json:"species" validate:"required"`
	Breed       string         `json:"breed" validate:"required"`
	Weight      float64        `json:"weight" validate:"required"`
	Gender      domain.EGender `json:"gender" validate:"required"`
}

type CreateResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func makePetCreateEndpoint(petService petservice.IPetService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(CreateRequest)
		if !ok {
			return nil, common.ErrCastRequest
		}

		if err := validator.New().Struct(req); err != nil {
			return nil, err
		}

		claims, ok := ctx.Value(kitjwt.JWTClaimsContextKey).(*jwt.StandardClaims)
		if !ok {
			return nil, common.ErrParsingClaims
		}

		userId := ownerdomain.UserId(claims.Subject)
		profile := domain.NewPetProfile(req.Name, req.DateOfBirth, req.Species, req.Breed, req.Weight, req.Gender)
		pet, err := petService.Create(ctx, userId, profile)
		if err != nil {
			return nil, err
		}

		return CreateResponse{
			string(pet.Id),
			pet.Profile.Name,
		}, nil
	}
}
