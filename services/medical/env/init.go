package env

import (
	"os"

	"github.com/dreadster3/pawcare/services/medical/services"
	"github.com/dreadster3/pawcare/shared/logger"
	sharedServices "github.com/dreadster3/pawcare/shared/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitEnvironment(db *mongo.Database) *Environment {
	logger := logger.Logger

	jwtSecret := os.Getenv("JWT_SECRET")

	servicesContainer := &ServiceContainer{
		auth:   sharedServices.NewAuthService(jwtSecret),
		record: services.NewRecordService(db),
	}

	return NewEnvironment(logger, servicesContainer)
}
