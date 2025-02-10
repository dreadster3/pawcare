package endpoint

import (
	"context"

	"github.com/dreadster3/pawcare/services/account/pet/domain"
	"github.com/dreadster3/pawcare/services/account/pet/service"
	"github.com/dreadster3/pawcare/services/auth"
	"github.com/dreadster3/pawcare/shared/common"
	"github.com/dreadster3/pawcare/shared/utils"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/golang-jwt/jwt/v4"
)

type GetAllResponse []GetResponse

func makeGetAllEndpoint(petService service.IPetService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		claims, ok := ctx.Value(kitjwt.JWTClaimsContextKey).(*jwt.StandardClaims)
		if !ok {
			return nil, common.ErrParsingClaims
		}

		userId := auth.UserId(claims.Subject)
		pets, err := petService.FindByUserId(ctx, userId)
		if err != nil {
			return nil, err
		}

		return utils.Map(pets, func(pet *domain.Pet) GetResponse {
			return GetResponse{
				string(pet.Id),
				pet.Profile.Name,
				pet.Profile.Species,
				pet.Profile.Breed,
			}
		}), nil
	}
}
