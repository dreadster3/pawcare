package endpoint

import (
	petservice "github.com/dreadster3/pawcare/services/account/internal/pet/service"
	"github.com/dreadster3/pawcare/shared/common"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type Set struct {
	CreateEndpoint  endpoint.Endpoint
	GetByIdEndpoint endpoint.Endpoint
	GetAllEndpoint  endpoint.Endpoint
}

func NewSet(viper viper.Viper, petService petservice.IPetService, logger log.Logger) Set {
	kf := common.JWTKeyFactory(viper)

	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = makePetCreateEndpoint(petService)
		createEndpoint = kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(createEndpoint)
	}

	var getByIdEndpoint endpoint.Endpoint
	{
		getByIdEndpoint = makeGetByIdEndpoint(petService)
		getByIdEndpoint = kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(getByIdEndpoint)
	}

	var getAllEndpoint endpoint.Endpoint
	{
		getAllEndpoint = makeGetAllEndpoint(petService)
		getAllEndpoint = kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(getAllEndpoint)
	}

	return Set{
		CreateEndpoint:  createEndpoint,
		GetByIdEndpoint: getByIdEndpoint,
		GetAllEndpoint:  getAllEndpoint,
	}
}
