package records

import (
	"github.com/dreadster3/pawcare/services/medical/env"
	sharedModels "github.com/dreadster3/pawcare/shared/models"
	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Create a new medical
// @Schemes
// @Description Creates a new medical record
// @Tags owner
// @Accept json
// @Param record body models.Record true "Record"
// @Produce json
// @Success 201 {object} models.Record
// @Failure 400 {object} sharedModels.ErrorResponse
// @Failure 401 {object} sharedModels.ErrorResponse
// @Failure 500 {object} sharedModels.ErrorResponse
// @Router /api/v1/records [post]
func Create(env *env.Environment, c *gin.Context) {
	c.JSON(501, sharedModels.NewErrorResponseString(c, "Not implemented"))
}
