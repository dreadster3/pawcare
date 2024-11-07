package env

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Environment struct {
	logger   *slog.Logger
	services *ServiceContainer
	viper    *viper.Viper
}

func NewEnvironment(logger *slog.Logger, viper *viper.Viper, serviceContainer *ServiceContainer) *Environment {
	return &Environment{
		logger:   logger,
		services: serviceContainer,
		viper:    viper,
	}
}

func (e *Environment) Logger() *slog.Logger {
	return e.logger
}

func (e *Environment) Services() *ServiceContainer {
	return e.services
}

func (e *Environment) WithRequestId(c *gin.Context) *Environment {
	return &Environment{
		logger:   e.logger.With("request_id", c.GetString("request_id")),
		services: e.services,
	}
}
