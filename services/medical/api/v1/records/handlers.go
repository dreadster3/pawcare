package records

import (
	"github.com/dreadster3/pawcare/services/medical/env"
	"github.com/dreadster3/pawcare/services/medical/models"
	"github.com/dreadster3/pawcare/services/medical/services"
	sharedModels "github.com/dreadster3/pawcare/shared/models"
	sharedServices "github.com/dreadster3/pawcare/shared/services"
	"github.com/gin-gonic/gin"
)

// Create godoc
// @Security JWT
// @Summary Create a new medical
// @Schemes
// @Description Creates a new medical record
// @Tags records
// @Accept json
// @Param body body models.Record true "Record"
// @Produce json
// @Success 201 {object} models.Record
// @Failure 400 {object} sharedModels.ErrorResponse
// @Failure 401 {object} sharedModels.ErrorResponse
// @Failure 404 {object} sharedModels.ErrorResponse
// @Failure 500 {object} sharedModels.ErrorResponse
// @Router /api/v1/records [post]
func Create(env *env.Environment, c *gin.Context) {
	var record models.Record
	if err := c.ShouldBindJSON(&record); err != nil {
		c.AbortWithStatusJSON(400, sharedModels.NewErrorResponse(c, err))
		return
	}

	userId := c.GetString("user_id")
	record.UserId = userId

	_, err := env.Services().Profile().Pet().FindById(record.PetId)
	if err != nil {
		if err == sharedServices.ErrNotFound {
			c.AbortWithStatusJSON(404, sharedModels.NewErrorResponseString(c, "Pet not found"))
			return
		}

		c.Error(err)
		c.AbortWithStatusJSON(500, sharedModels.NewInternalErrorResponse(c))
		return
	}

	newRecord, err := env.Services().Record().Create(record)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(500, sharedModels.NewErrorResponseString(c, "Failed to create record"))
		return
	}

	c.JSON(200, newRecord)
}

// GetByPetId godoc
// @Security JWT
// @Summary Get Records by Pet Id
// @Description Get All Records that belong to a Pet
// @Tags records
// @Produce json
// @Param id path string true "Pet Id"
// @Success 200 {object} []models.Record
// @Failure 401 {object} sharedModels.ErrorResponse
// @Failure 500 {object} sharedModels.ErrorResponse
// @Router /api/v1/pets/:id/records [get]
func GetByPetId(env *env.Environment, c *gin.Context) {
	petId := c.Param("id")
	userId := c.GetString("user_id")

	records, err := env.Services().Record().FindByUserIdAndPetId(userId, petId)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(500, sharedModels.NewErrorResponseString(c, "Failed to get records"))
		return
	}

	c.JSON(200, records)
}

// GetById godoc
// @Security JWT
// @Summary Get Record
// @Description Get Record with the specified Id
// @Tags records
// @Produce json
// @Param id path string true "Record Id"
// @Success 200 {object} models.Record
// @Failure 401 {object} sharedModels.ErrorResponse
// @Failure 404 {object} sharedModels.ErrorResponse
// @Failure 500 {object} sharedModels.ErrorResponse
// @Router /api/v1/records/:id [get]
func GetById(env *env.Environment, c *gin.Context) {
	id := c.Param("id")
	userId := c.GetString("user_id")

	record, err := env.Services().Record().FindByUserIdAndId(userId, id)
	if err != nil {
		if err == services.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, sharedModels.NewErrorResponse(c, err))
			return
		}

		c.Error(err)
		c.AbortWithStatusJSON(500, sharedModels.NewErrorResponseString(c, "Failed to get record"))
		return
	}

	c.JSON(200, record)
}
