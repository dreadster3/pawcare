package healthcheck

import (
	"github.com/dreadster3/pawcare/services/medical/env"
	"github.com/dreadster3/pawcare/services/medical/services"
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
func HealthcheckGet(env *env.Environment, c *gin.Context) {
	report := services.HealthcheckService.Run(c.Request.Context())

	c.JSON(200, report)
}
