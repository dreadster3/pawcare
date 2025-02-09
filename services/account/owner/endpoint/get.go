package endpoint

import (
	"context"
	"time"

	"github.com/dreadster3/pawcare/services/account/owner/service"
	"github.com/dreadster3/pawcare/services/auth"
	"github.com/dreadster3/pawcare/shared/common"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/golang-jwt/jwt/v4"
)

type GetResponse struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

func makeGetEndpoint(ownerService service.IOwnerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		claims, ok := ctx.Value(kitjwt.JWTClaimsContextKey).(*jwt.StandardClaims)
		if !ok {
			return nil, common.ErrParsingClaims
		}

		userId := auth.UserId(claims.Subject)
		owner, err := ownerService.FindByUserId(ctx, userId)
		if err != nil {
			return nil, err
		}

		return GetResponse{string(owner.Id), owner.Profile.Name, owner.Profile.DateOfBirth}, nil
	}
}
