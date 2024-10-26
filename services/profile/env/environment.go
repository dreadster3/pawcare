package env

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Environment struct {
	logger   *slog.Logger
	services *ServiceContainer
}

func NewEnvironment(logger *slog.Logger, serviceContainer *ServiceContainer) *Environment {
	return &Environment{
		logger:   logger,
		services: serviceContainer,
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
