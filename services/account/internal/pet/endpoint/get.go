package endpoint

import (
	"context"

	ownerdomain "github.com/dreadster3/pawcare/services/account/internal/owner/domain"
	ownerservice "github.com/dreadster3/pawcare/services/account/internal/owner/service"
	"github.com/dreadster3/pawcare/services/account/internal/pet/domain"
	petservice "github.com/dreadster3/pawcare/services/account/internal/pet/service"
	"github.com/dreadster3/pawcare/shared/common"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type GetRequest struct {
	Id string `json:"id" validate:"required"`
}

type GetResponse struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Species string `json:"species"`
	Breed   string `json:"breed"`
}

func makeGetEndpoint(ownerService ownerservice.IOwnerService, petService petservice.IPetService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(GetRequest)
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
		owner, err := ownerService.FindByUserId(ctx, userId)
		if err != nil {
			return nil, err
		}

		pet, err := petService.FindById(ctx, domain.PetId(req.Id))
		if err != nil {
			return nil, err
		}

		if owner.Id != pet.OwnerId {
			return nil, common.ErrUnauthorized
		}

		return GetResponse{
			string(pet.Id),
			pet.Profile.Name,
			pet.Profile.Species,
			pet.Profile.Breed,
		}, nil
	}
}
