package env

import (
	"os"

	"github.com/dreadster3/pawcare/services/profile/repository/mongodb"
	"github.com/dreadster3/pawcare/services/profile/services"
	"github.com/dreadster3/pawcare/shared/logger"
	sharedServices "github.com/dreadster3/pawcare/shared/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitEnvironment(db *mongo.Database) *Environment {
	logger := logger.Logger

	petRepository := mongodb.NewPetRepository(db)
	ownerRepository := mongodb.NewOwnerRepository(db)

	jwtSecret := os.Getenv("JWT_SECRET")

	servicesContainer := &ServiceContainer{
		auth:  sharedServices.NewAuthService(jwtSecret),
		pet:   services.NewPetService(petRepository),
		owner: services.NewOwnerService(ownerRepository),
	}

	return NewEnvironment(logger, servicesContainer)
}
