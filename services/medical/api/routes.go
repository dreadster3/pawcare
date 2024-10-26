package api

import (
	"github.com/dreadster3/pawcare/services/medical/api/healthcheck"
	"github.com/dreadster3/pawcare/services/medical/api/v1/records"
	"github.com/dreadster3/pawcare/services/medical/env"
	"github.com/dreadster3/pawcare/shared/handlers"
	"github.com/dreadster3/pawcare/shared/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(environment *env.Environment, router *gin.RouterGroup) {
	apiGroup := router.Group("/api")

	healthcheckGroup := apiGroup.Group("/healthcheck")
	healthcheckGroup.GET("", handlers.WrapperEnv(environment, healthcheck.HealthcheckGet))

	v1Group := apiGroup.Group("/v1", middleware.JwtAuth(environment))

	{
		recordsGroup := v1Group.Group("/records")
		recordsGroup.POST("", handlers.WrapperEnv(environment, records.Create))
	}
}
