package api

import (
	"github.com/dreadster3/pawcare/services/profile/api/healthcheck"
	"github.com/dreadster3/pawcare/services/profile/api/v1/profiles/owner"
	"github.com/dreadster3/pawcare/services/profile/api/v1/profiles/pet"
	"github.com/dreadster3/pawcare/services/profile/env"
	"github.com/dreadster3/pawcare/services/profile/middleware"
	"github.com/gin-gonic/gin"
)

type HandlerFuncWithEnv func(env *env.Environment, c *gin.Context)

func EnvWrapper(env *env.Environment, handler HandlerFuncWithEnv) gin.HandlerFunc {
	return func(c *gin.Context) {
		env := env.WithRequestId(c)
		handler(env, c)
	}
}

func RegisterRoutes(router *gin.RouterGroup) {
	environment := env.InitEnvironment()

	apiGroup := router.Group("/api")

	healthcheckGroup := apiGroup.Group("/healthcheck")
	healthcheckGroup.GET("", EnvWrapper(environment, healthcheck.HealthcheckGet))

	petGroup := apiGroup.Group("/v1")
	profileGroup := petGroup.Group("/profiles", middleware.JwtAuth(environment))

	{
		petProfileGroup := profileGroup.Group("/pets")
		petProfileGroup.POST("", EnvWrapper(environment, pet.Create))
		petProfileGroup.GET("", EnvWrapper(environment, pet.GetAll))
		petProfileGroup.GET("/:id", EnvWrapper(environment, pet.GetById))
	}

	{
		ownerProfileGroup := profileGroup.Group("/owners")
		ownerProfileGroup.POST("", EnvWrapper(environment, owner.Create))
		ownerProfileGroup.GET("", EnvWrapper(environment, owner.Get))
	}
}
