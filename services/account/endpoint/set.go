package endpoint

import (
	ownerservice "github.com/dreadster3/pawcare/services/account/service/owner"
	petservice "github.com/dreadster3/pawcare/services/account/service/pet"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/golang-jwt/jwt/v4"
)

type Set struct {
	OwnerCreateEndpoint endpoint.Endpoint
	PetCreateEndpoint   endpoint.Endpoint
}

func NewSet(ownerService ownerservice.IOwnerService, petService petservice.IPetService, logger log.Logger) Set {
	kf := func(token *jwt.Token) (interface{}, error) { return []byte("SuperSecret"), nil }

	var ownerCreateEndpoint endpoint.Endpoint
	{
		ownerCreateEndpoint = makeOwnerCreateEndpoint(ownerService)
		ownerCreateEndpoint = kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(ownerCreateEndpoint)
	}

	var petCreateEndpoint endpoint.Endpoint
	{
		petCreateEndpoint = makePetCreateEndpoint(ownerService, petService)
		petCreateEndpoint = kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(petCreateEndpoint)
	}

	return Set{
		OwnerCreateEndpoint: ownerCreateEndpoint,
		PetCreateEndpoint:   petCreateEndpoint,
	}
}
