package api

import (
	"github.com/dreadster3/pawcare/services/profile/api/v1/profiles/owner"
	"github.com/dreadster3/pawcare/services/profile/api/v1/profiles/pet"
	"github.com/dreadster3/pawcare/services/profile/env"
	"github.com/dreadster3/pawcare/shared/handlers"
	"github.com/dreadster3/pawcare/shared/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

func EnvironmentFactory(viper *viper.Viper, db *mongo.Database) handlers.EnvFactory[*env.ServiceContainer, *env.Environment] {
	return func(c *gin.Context) (*env.Environment, error) {
		return env.InitEnvironment(viper, db), nil
	}
}

func RegisterRoutes(viper *viper.Viper, db *mongo.Database, router *gin.RouterGroup) {
	apiGroup := router.Group("/api")

	factory := EnvironmentFactory(viper, db)
	healthcheckGroup := apiGroup.Group("/healthcheck")
	healthcheckGroup.GET("", handlers.WrapperEnv(factory, handlers.HealthcheckHandler))

	petGroup := apiGroup.Group("/v1")
	profileGroup := petGroup.Group("/profiles", middleware.JwtAuth(factory))

	{
		petProfileGroup := profileGroup.Group("/pets")
		petProfileGroup.POST("", handlers.WrapperEnv(factory, pet.Create))
		petProfileGroup.GET("", handlers.WrapperEnv(factory, pet.GetAll))
		petProfileGroup.GET("/:id", handlers.WrapperEnv(factory, pet.GetById))
	}

	{
		ownerProfileGroup := profileGroup.Group("/owners")
		ownerProfileGroup.POST("", handlers.WrapperEnv(factory, owner.Create))
		ownerProfileGroup.GET("", handlers.WrapperEnv(factory, owner.Get))
	}
}
