package api

import (
	"github.com/dreadster3/pawcare/services/profile/api/healthcheck"
	"github.com/dreadster3/pawcare/services/profile/api/v1/profiles/owner"
	"github.com/dreadster3/pawcare/services/profile/api/v1/profiles/pet"
	"github.com/dreadster3/pawcare/services/profile/env"
	"github.com/dreadster3/pawcare/shared/handlers"
	"github.com/dreadster3/pawcare/shared/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(environment *env.Environment, router *gin.RouterGroup) {
	apiGroup := router.Group("/api")

	healthcheckGroup := apiGroup.Group("/healthcheck")
	healthcheckGroup.GET("", handlers.WrapperEnv(environment, healthcheck.HealthcheckGet))

	petGroup := apiGroup.Group("/v1")
	profileGroup := petGroup.Group("/profiles", middleware.JwtAuth(environment))

	{
		petProfileGroup := profileGroup.Group("/pets")
		petProfileGroup.POST("", handlers.WrapperEnv(environment, pet.Create))
		petProfileGroup.GET("", handlers.WrapperEnv(environment, pet.GetAll))
		petProfileGroup.GET("/:id", handlers.WrapperEnv(environment, pet.GetById))
	}

	{
		ownerProfileGroup := profileGroup.Group("/owners")
		ownerProfileGroup.POST("", handlers.WrapperEnv(environment, owner.Create))
		ownerProfileGroup.GET("", handlers.WrapperEnv(environment, owner.Get))
	}
}
