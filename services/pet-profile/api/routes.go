package api

import (
	"github.com/dreadster3/pawcare/services/pet-profile/api/healthcheck"
	"github.com/dreadster3/pawcare/services/pet-profile/api/v1/profile"
	"github.com/dreadster3/pawcare/services/pet-profile/env"
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
	{
		profileGroup := petGroup.Group("/profiles")
		profileGroup.POST("", EnvWrapper(environment, profile.Create))
		profileGroup.GET("", EnvWrapper(environment, profile.GetAll))
		profileGroup.GET("/:id", EnvWrapper(environment, profile.GetByID))
	}
}
