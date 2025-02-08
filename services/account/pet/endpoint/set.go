package endpoint

import (
	ownerservice "github.com/dreadster3/pawcare/services/account/owner/service"
	petservice "github.com/dreadster3/pawcare/services/account/pet/service"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/golang-jwt/jwt/v4"
)

type Set struct {
	CreateEndpoint endpoint.Endpoint
}

func NewSet(ownerService ownerservice.IOwnerService, petService petservice.IPetService, logger log.Logger) Set {
	kf := func(token *jwt.Token) (interface{}, error) { return []byte("SuperSecret"), nil }

	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = makePetCreateEndpoint(ownerService, petService)
		createEndpoint = kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(createEndpoint)
	}

	return Set{
		CreateEndpoint: createEndpoint,
	}
}
