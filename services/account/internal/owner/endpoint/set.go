package endpoint

import (
	"github.com/dreadster3/pawcare/services/account/internal/owner/service"
	"github.com/dreadster3/pawcare/shared/oauth"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/spf13/viper"
)

type Set struct {
	CreateEndpoint endpoint.Endpoint
	GetEndpoint    endpoint.Endpoint
}

func NewSet(viper viper.Viper, ownerService service.IOwnerService, logger log.Logger) Set {

	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = makeCreateEndpoint(ownerService)
		createEndpoint = oauth.NewConfiguredIntrospectionMiddleware(viper)(createEndpoint)
	}

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = makeGetEndpoint(ownerService)
		getEndpoint = oauth.NewConfiguredIntrospectionMiddleware(viper)(getEndpoint)
	}

	return Set{
		CreateEndpoint: createEndpoint,
		GetEndpoint:    getEndpoint,
	}
}
