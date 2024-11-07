package api

import (
	"github.com/dreadster3/pawcare/services/medical/api/v1/records"
	"github.com/dreadster3/pawcare/services/medical/env"
	"github.com/dreadster3/pawcare/shared/handlers"
	"github.com/dreadster3/pawcare/shared/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

func EnvironmentFactory(viper *viper.Viper, db *mongo.Database) handlers.EnvFactory[*env.ServiceContainer, *env.Environment] {
	return func(c *gin.Context) (*env.Environment, error) {
		return env.InitEnvironment(c, viper, db)
	}
}

func RegisterRoutes(viper *viper.Viper, db *mongo.Database, router *gin.RouterGroup) {
	apiGroup := router.Group("/api")
	factory := EnvironmentFactory(viper, db)
	healthcheckGroup := apiGroup.Group("/healthcheck")
	healthcheckGroup.GET("", handlers.WrapperEnv(factory, handlers.HealthcheckHandler))

	v1Group := apiGroup.Group("/v1", middleware.JwtAuth(factory))

	{
		recordsGroup := v1Group.Group("/records")
		recordsGroup.POST("", handlers.WrapperEnv(factory, records.Create))
		recordsGroup.GET(":id", handlers.WrapperEnv(factory, records.GetById))
	}

	{
		petsGroup := v1Group.Group("/pets")
		petsGroup.GET("/:id/records", handlers.WrapperEnv(factory, records.GetByPetId))
	}
}
