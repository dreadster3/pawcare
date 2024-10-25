package env

import (
	"os"

	"github.com/dreadster3/pawcare/services/profile/logger"
	"github.com/dreadster3/pawcare/services/profile/repository/mongodb"
	"github.com/dreadster3/pawcare/services/profile/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitEnvironment(db *mongo.Database) *Environment {
	logger := logger.Logger

	petRepository := mongodb.NewPetRepository(db)
	ownerRepository := mongodb.NewOwnerRepository(db)

	jwtSecret := os.Getenv("JWT_SECRET")

	servicesContainer := &ServiceContainer{
		Auth:  services.NewAuthService(jwtSecret),
		Pet:   services.NewPetService(petRepository),
		Owner: services.NewOwnerService(ownerRepository),
	}

	return NewEnvironment(logger, servicesContainer)
}
