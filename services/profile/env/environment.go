package env

import (
	"log/slog"

	"github.com/dreadster3/pawcare/services/profile/services"
	"github.com/gin-gonic/gin"
)

type ServiceContainer struct {
	Auth  *services.AuthService
	Pet   *services.PetService
	Owner *services.OwnerService
}

type Environment struct {
	Logger   *slog.Logger
	Services *ServiceContainer
}

func NewEnvironment(logger *slog.Logger, serviceContainer *ServiceContainer) *Environment {
	return &Environment{
		Logger:   logger,
		Services: serviceContainer,
	}
}

func (e *Environment) WithRequestId(c *gin.Context) *Environment {
	return &Environment{
		Logger:   e.Logger.With("request_id", c.GetString("request_id")),
		Services: e.Services,
	}
}
