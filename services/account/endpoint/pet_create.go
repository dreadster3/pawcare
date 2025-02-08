package endpoint

import (
	"context"
	"time"

	ownerservice "github.com/dreadster3/pawcare/services/account/service/owner"
	petservice "github.com/dreadster3/pawcare/services/account/service/pet"
	"github.com/dreadster3/pawcare/services/account/valueobjects"
	"github.com/dreadster3/pawcare/services/auth"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type PetCreateRequest struct {
	Name        string               `json:"name" validate:"required"`
	DateOfBirth time.Time            `json:"date_of_birth" validate:"required"`
	Species     string               `json:"species" validate:"required"`
	Breed       string               `json:"breed" validate:"required"`
	Weight      float64              `json:"weight" validate:"required"`
	Gender      valueobjects.EGender `json:"gender" validate:"required"`
}

type PetCreateResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func makePetCreateEndpoint(ownerService ownerservice.IOwnerService, petService petservice.IPetService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(PetCreateRequest)
		if !ok {
			return nil, ErrCastRequest
		}

		if err := validator.New().Struct(req); err != nil {
			return nil, err
		}

		claims, ok := ctx.Value(kitjwt.JWTClaimsContextKey).(*jwt.StandardClaims)
		if !ok {
			return nil, ErrParsingClaims
		}

		userId := auth.UserId(claims.Subject)
		owner, err := ownerService.FindByUserId(ctx, userId)
		if err != nil {
			return nil, err
		}

		profile := valueobjects.NewPetProfile(req.Name, req.DateOfBirth, req.Species, req.Breed, req.Weight, req.Gender)
		pet, err := petService.Create(ctx, owner.Id, profile)
		if err != nil {
			return nil, err
		}

		return PetCreateResponse{
			string(pet.Id),
			pet.Profile.Name,
		}, nil
	}
}
