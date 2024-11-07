package env

import (
	"github.com/dreadster3/pawcare/services/profile/repository/mongodb"
	"github.com/dreadster3/pawcare/services/profile/services"
	"github.com/dreadster3/pawcare/shared/logger"
	sharedServices "github.com/dreadster3/pawcare/shared/services"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitEnvironment(viper *viper.Viper, db *mongo.Database) *Environment {
	logger := logger.Logger

	petRepository := mongodb.NewPetRepository(db)
	ownerRepository := mongodb.NewOwnerRepository(db)

	secret := viper.GetString("JWT_SECRET")

	servicesContainer := &ServiceContainer{
		auth:  sharedServices.NewAuthService(secret),
		pet:   services.NewPetService(petRepository),
		owner: services.NewOwnerService(ownerRepository),

		healthcheck: sharedServices.NewHealthcheckService(db),
	}

	return NewEnvironment(logger, servicesContainer)
}
