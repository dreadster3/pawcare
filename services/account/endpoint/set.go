package endpoint

import (
	"github.com/dreadster3/pawcare/services/account/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

type Set struct {
	CreateAccountEndpoint endpoint.Endpoint
}

func NewSet(svc service.ProfileService, logger log.Logger) Set {
	var createAccountEndpoint endpoint.Endpoint
	{
		createAccountEndpoint = makeCreateOwnerEndpoint(svc)
	}

	return Set{
		CreateAccountEndpoint: createAccountEndpoint,
	}
}
