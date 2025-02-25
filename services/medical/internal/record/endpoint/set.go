package endpoint

import (
	"github.com/dreadster3/pawcare/services/medical/internal/record/service"
	"github.com/dreadster3/pawcare/shared/oauth"
	"github.com/go-kit/kit/endpoint"
	"github.com/spf13/viper"
)

type Set struct {
	CreateEndpoint     endpoint.Endpoint
	GetByPetIdEndpoint endpoint.Endpoint
	GetByIdEndpoint    endpoint.Endpoint
}

func NewSet(viper viper.Viper, recordService service.IRecordService) Set {
	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = makeCreateEndpoint(recordService)
		createEndpoint = oauth.NewConfiguredIntrospectionMiddleware(viper)(createEndpoint)
	}

	var getByPetIdEndpoint endpoint.Endpoint
	{
		getByPetIdEndpoint = makeGetByPetIdEndpoint(recordService)
		getByPetIdEndpoint = oauth.NewConfiguredIntrospectionMiddleware(viper)(getByPetIdEndpoint)
	}

	var getByIdEndpoint endpoint.Endpoint
	{
		getByIdEndpoint = makeGetByIdEndpoint(recordService)
		getByIdEndpoint = oauth.NewConfiguredIntrospectionMiddleware(viper)(getByIdEndpoint)
	}

	return Set{
		CreateEndpoint:     createEndpoint,
		GetByPetIdEndpoint: getByPetIdEndpoint,
		GetByIdEndpoint:    getByIdEndpoint,
	}
}
