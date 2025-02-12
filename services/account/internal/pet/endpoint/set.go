package endpoint

import (
	ownerservice "github.com/dreadster3/pawcare/services/account/internal/owner/service"
	petservice "github.com/dreadster3/pawcare/services/account/internal/pet/service"
	"github.com/dreadster3/pawcare/shared/common"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type Set struct {
	CreateEndpoint endpoint.Endpoint
	GetEndpoint    endpoint.Endpoint
	GetAllEndpoint endpoint.Endpoint
}

func NewSet(viper viper.Viper, ownerService ownerservice.IOwnerService, petService petservice.IPetService, logger log.Logger) Set {
	kf := common.JWTKeyFactory(viper)

	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = makePetCreateEndpoint(petService)
		createEndpoint = kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(createEndpoint)
	}

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = makeGetEndpoint(ownerService, petService)
		getEndpoint = kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(getEndpoint)
	}

	var getAllEndpoint endpoint.Endpoint
	{
		getAllEndpoint = makeGetAllEndpoint(petService)
		getAllEndpoint = kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(getAllEndpoint)
	}

	return Set{
		CreateEndpoint: createEndpoint,
		GetEndpoint:    getEndpoint,
		GetAllEndpoint: getAllEndpoint,
	}
}
