package handlers

import (
	"github.com/dreadster3/pawcare/shared/env"
	"github.com/gin-gonic/gin"
)

// HealthcheckGet godoc
// @Summary Endpoint to check if the service is healthy
// @Schemes
// @Description Checks if the internal services used by the service are healthy
// @Tags healthcheck
// @Produce json
// @Success 200 {object} healthcheck.JSONHealthcheckReport
// @Router /api/healthcheck [get]
func HealthcheckHandler[T env.IServiceContainer, E env.IEnvironment[T]](env E, c *gin.Context) {
	report := env.Services().Healthcheck().Run(c)

	c.JSON(200, report)
}
