package env

import (
	"github.com/dreadster3/pawcare/services/pet-profile/logger"
	"github.com/dreadster3/pawcare/services/pet-profile/repository/mongodb"
	"github.com/dreadster3/pawcare/services/pet-profile/services"
)

func InitEnvironment() *Environment {
	logger := logger.Logger

	petProfileRepository := mongodb.NewPetProfileRepository()

	servicesContainer := &ServiceContainer{
		PetProfile: services.NewPetProfileService(petProfileRepository),
	}

	return NewEnvironment(logger, servicesContainer)
}
