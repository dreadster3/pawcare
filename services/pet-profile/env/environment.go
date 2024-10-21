package env

import (
	"context"
	"log/slog"

	"github.com/dreadster3/pawcare/services/pet-profile/db"
	"github.com/dreadster3/pawcare/services/pet-profile/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceContainer struct {
	PetProfile *services.PetProfileService
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

func (e Environment) Database() (*mongo.Database, db.DbCloseFunc, error) {
	ctx := context.Background()

	e.Logger.Debug("Connecting to database...")

	database, disconnect, err := db.ConnectDB(ctx)
	new_disconnect := func(ctx context.Context) error {
		e.Logger.Debug("Disconnecting from database...")
		return disconnect(ctx)
	}

	return database, new_disconnect, err
}

func (e *Environment) WithRequestId(c *gin.Context) *Environment {
	return &Environment{
		Logger:   e.Logger.With("request_id", c.GetString("request_id")),
		Services: e.Services,
	}
}
