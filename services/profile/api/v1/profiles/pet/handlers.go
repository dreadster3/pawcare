package pet

import (
	"github.com/dreadster3/pawcare/services/profile/env"
	"github.com/dreadster3/pawcare/services/profile/models"
	"github.com/dreadster3/pawcare/services/profile/services"
	sharedModels "github.com/dreadster3/pawcare/shared/models"
	"github.com/gin-gonic/gin"
)

// Create godoc
// @Security JWT
// @Summary Create a new pet profile
// @Description Creates a new pet profile
// @Tags pet
// @Accept json
// @Param body body models.Pet true "Pet Profile"
// @Produce json
// @Success 201 {object} models.Pet
// @Failure 400 {object} sharedModels.ErrorResponse
// @Failure 401 {object} sharedModels.ErrorResponse
// @Failure 500 {object} sharedModels.ErrorResponse
// @Router /api/v1/profiles/pets [post]
func Create(env *env.Environment, c *gin.Context) {
	var request models.Pet
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(400, sharedModels.NewErrorResponse(c, err))
		return
	}

	userId := c.GetString("user_id")
	owner, err := env.Services().Owner().FindByUserId(userId)
	if err != nil {
		if err == services.ErrProfileNotFound {
			c.Error(err)
			c.AbortWithStatusJSON(404, sharedModels.NewErrorResponseString(c, "Owner not found"))
			return
		}

		c.Error(err)
		c.AbortWithStatusJSON(500, sharedModels.NewErrorResponseString(c, "Internal server error"))
		return
	}

	request.OwnerId = owner.Id.Hex()
	result, err := env.Services().Pet().Create(request)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(500, sharedModels.NewErrorResponseString(c, "Internal server error"))
		return
	}

	c.JSON(201, result)
}

// GetAll godoc
// @Security JWT
// @Summary Get all pet profiles
// @Description Get all pet profiles
// @Tags pet
// @Produce json
// @Success 200 {array} models.Pet
// @Failure 401 {object} sharedModels.ErrorResponse
// @Failure 500 {object} sharedModels.ErrorResponse
// @Router /api/v1/profiles/pets [get]
func GetAll(env *env.Environment, c *gin.Context) {

	userId := c.GetString("user_id")
	owner, err := env.Services().Owner().FindByUserId(userId)
	if err != nil {
		if err == services.ErrProfileNotFound {
			c.Error(err)
			c.AbortWithStatusJSON(404, sharedModels.NewErrorResponseString(c, "Owner not found"))
			return
		}

		c.Error(err)
		c.AbortWithStatusJSON(500, sharedModels.NewErrorResponseString(c, "Internal server error"))
		return
	}

	result, err := env.Services().Pet().FindByOwnerId(owner.Id.Hex())
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(500, sharedModels.NewErrorResponseString(c, "Internal server error"))
		return
	}

	c.JSON(200, result)
}

// GetById godoc
// @Security JWT
// @Summary Get pet profile by ID
// @Description Get pet profile by ID
// @Tags pet
// @Param id path string true "Pet Profile ID"
// @Produce json
// @Success 200 {object} models.Pet
// @Failure 401 {object} sharedModels.ErrorResponse
// @Failure 404 {object} sharedModels.ErrorResponse
// @Failure 500 {object} sharedModels.ErrorResponse
// @Router /api/v1/profiles/pets/{id} [get]
func GetById(env *env.Environment, c *gin.Context) {
	id := c.Param("id")

	userId := c.GetString("user_id")
	owner, err := env.Services().Owner().FindByUserId(userId)
	if err != nil {
		if err == services.ErrProfileNotFound {
			c.Error(err)
			c.AbortWithStatusJSON(404, sharedModels.NewErrorResponseString(c, "Owner not found"))
			return
		}

		c.Error(err)
		c.AbortWithStatusJSON(500, sharedModels.NewErrorResponseString(c, "Internal server error"))
		return
	}

	result, err := env.Services().Pet().FindByIdAndOwnerId(id, owner.Id.Hex())
	if err != nil {
		if err == services.ErrInvalidId {
			c.Error(err)
			c.AbortWithStatusJSON(400, sharedModels.NewErrorResponseString(c, "Invalid ID"))
			return
		}

		if err == services.ErrProfileNotFound {
			c.Error(err)
			c.AbortWithStatusJSON(404, sharedModels.NewErrorResponseString(c, "Profile not found"))
			return
		}

		c.Error(err)
		c.AbortWithStatusJSON(500, sharedModels.NewErrorResponseString(c, "Internal server error"))
		return
	}

	c.JSON(200, result)
}
