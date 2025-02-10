package endpoint

import (
	"github.com/dreadster3/pawcare/services/account/owner/service"
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
}

func NewSet(viper viper.Viper, ownerService service.IOwnerService, logger log.Logger) Set {
	kf := common.JWTKeyFactory(viper)

	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = makeCreateEndpoint(ownerService)
		createEndpoint = kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(createEndpoint)
	}

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = makeGetEndpoint(ownerService)
		getEndpoint = kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(getEndpoint)
	}

	return Set{
		CreateEndpoint: createEndpoint,
		GetEndpoint:    getEndpoint,
	}
}
