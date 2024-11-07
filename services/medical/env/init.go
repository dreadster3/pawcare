package env

import (
	"github.com/dreadster3/pawcare/services/medical/repository/mongodb"
	"github.com/dreadster3/pawcare/services/medical/services"
	"github.com/dreadster3/pawcare/shared/config"
	"github.com/dreadster3/pawcare/shared/logger"
	sharedServices "github.com/dreadster3/pawcare/shared/services"
	"github.com/dreadster3/pawcare/shared/services/http/profiles"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitEnvironment(ctx *gin.Context, viper *viper.Viper, db *mongo.Database) (*Environment, error) {
	logger := logger.Logger

	recordRepository := mongodb.NewRecordRepository(db)
	jwtSecret := viper.GetString("jwt_secret")

	profileConfig, err := config.LoadServiceConfig(viper, "profile")
	if err != nil {
		return nil, err
	}

	healthcheckService := sharedServices.NewHealthcheckService(db)
	healthcheckService.RegisterServiceCheck("profile", profileConfig)

	servicesContainer := &ServiceContainer{
		auth:        sharedServices.NewAuthService(jwtSecret),
		record:      services.NewRecordService(recordRepository),
		profile:     profiles.NewProfileService(profileConfig, ctx),
		healthcheck: healthcheckService,
	}

	return NewEnvironment(logger, viper, servicesContainer), nil
}
