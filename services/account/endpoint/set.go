package endpoint

import (
	"github.com/dreadster3/pawcare/services/account/service"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/golang-jwt/jwt/v4"
)

type Set struct {
	CreateAccountEndpoint endpoint.Endpoint
	GetAccountEndpoint    endpoint.Endpoint
}

func NewSet(profileService service.IAccountService, logger log.Logger) Set {
	kf := func(token *jwt.Token) (interface{}, error) { return []byte("SuperSecret"), nil }

	var createAccountEndpoint endpoint.Endpoint
	{
		createAccountEndpoint = makeCreateAccountEndpoint(profileService)
		createAccountEndpoint = kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(createAccountEndpoint)
	}

	var getAccountEndpoint endpoint.Endpoint
	{
		getAccountEndpoint = makeGetAccountEndpoint(profileService)
		getAccountEndpoint = kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(getAccountEndpoint)
	}

	return Set{
		CreateAccountEndpoint: createAccountEndpoint,
		GetAccountEndpoint:    getAccountEndpoint,
	}
}
