package env

import (
	"log/slog"

	"github.com/dreadster3/pawcare/shared/services"
)

type IServiceContainer interface {
	Auth() services.IAuthService
	Healthcheck() services.IHealthcheckService
}

type IEnvironment[T IServiceContainer] interface {
	Logger() *slog.Logger
	Services() T
}
