package endpoint

import (
	petservice "github.com/dreadster3/pawcare/services/account/internal/pet/service"
	"github.com/dreadster3/pawcare/shared/oauth"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/spf13/viper"
)

type Set struct {
	CreateEndpoint  endpoint.Endpoint
	GetByIdEndpoint endpoint.Endpoint
	GetAllEndpoint  endpoint.Endpoint
}

func NewSet(viper viper.Viper, petService petservice.IPetService, logger log.Logger) Set {
	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = makePetCreateEndpoint(petService)
		createEndpoint = oauth.NewConfiguredIntrospectionMiddleware(viper)(createEndpoint)
	}

	var getByIdEndpoint endpoint.Endpoint
	{
		getByIdEndpoint = makeGetByIdEndpoint(petService)
		getByIdEndpoint = oauth.NewConfiguredIntrospectionMiddleware(viper)(getByIdEndpoint)
	}

	var getAllEndpoint endpoint.Endpoint
	{
		getAllEndpoint = makeGetAllEndpoint(petService)
		getAllEndpoint = oauth.NewConfiguredIntrospectionMiddleware(viper)(getAllEndpoint)
	}

	return Set{
		CreateEndpoint:  createEndpoint,
		GetByIdEndpoint: getByIdEndpoint,
		GetAllEndpoint:  getAllEndpoint,
	}
}
