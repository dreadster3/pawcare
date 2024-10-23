package env

import (
	"os"

	"github.com/dreadster3/pawcare/services/profile/logger"
	"github.com/dreadster3/pawcare/services/profile/repository/mongodb"
	"github.com/dreadster3/pawcare/services/profile/services"
)

func InitEnvironment() *Environment {
	logger := logger.Logger

	petRepository := mongodb.NewPetRepository()
	ownerRepository := mongodb.NewOwnerRepository()

	jwtSecret := os.Getenv("JWT_SECRET")

	servicesContainer := &ServiceContainer{
		Auth:  services.NewAuthService(jwtSecret),
		Pet:   services.NewPetService(petRepository),
		Owner: services.NewOwnerService(ownerRepository),
	}

	return NewEnvironment(logger, servicesContainer)
}
