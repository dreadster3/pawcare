package endpoint

import (
	ownerservice "github.com/dreadster3/pawcare/services/account/service/owner_service"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/golang-jwt/jwt/v4"
)

type Set struct {
	CreateOwnerEndpoint endpoint.Endpoint
}

func NewSet(ownerService ownerservice.IOwnerService, logger log.Logger) Set {
	kf := func(token *jwt.Token) (interface{}, error) { return []byte("SuperSecret"), nil }

	var createAccountEndpoint endpoint.Endpoint
	{
		createAccountEndpoint = makeCreateOwnerEndpoint(ownerService)
		createAccountEndpoint = kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(createAccountEndpoint)
	}

	return Set{
		CreateOwnerEndpoint: createAccountEndpoint,
	}
}
