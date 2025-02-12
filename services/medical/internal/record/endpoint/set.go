package endpoint

import (
	"github.com/dreadster3/pawcare/services/medical/internal/record/service"
	"github.com/dreadster3/pawcare/shared/common"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type Set struct {
	CreateEndpoint     endpoint.Endpoint
	GetByPetIdEndpoint endpoint.Endpoint
}

func NewSet(viper viper.Viper, recordService service.IRecordService) Set {
	kf := common.JWTKeyFactory(viper)

	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = makeCreateEndpoint(recordService)
		createEndpoint = kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(createEndpoint)
	}

	var getByPetIdEndpoint endpoint.Endpoint
	{
		getByPetIdEndpoint = makeGetByPetIdEndpoint(recordService)
		getByPetIdEndpoint = kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(getByPetIdEndpoint)
	}

	return Set{
		CreateEndpoint:     createEndpoint,
		GetByPetIdEndpoint: getByPetIdEndpoint,
	}
}
