package owner

import (
	"github.com/dreadster3/pawcare/services/profile/env"
	"github.com/dreadster3/pawcare/services/profile/models"
	"github.com/dreadster3/pawcare/services/profile/services"
	sharedModels "github.com/dreadster3/pawcare/shared/models"
	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Create a new owner profile
// @Schemes
// @Description Creates a new owner profile
// @Tags owner
// @Accept json
// @Param owner_profile body models.Owner true "Owner Profile"
// @Produce json
// @Success 201 {object} models.Owner
// @Failure 400 {object} sharedModels.ErrorResponse
// @Failure 401 {object} sharedModels.ErrorResponse
// @Failure 500 {object} sharedModels.ErrorResponse
// @Router /api/v1/profiles/owners [post]
func Create(env *env.Environment, c *gin.Context) {
	var request models.Owner
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(400, sharedModels.NewErrorResponse(c, err))
		return
	}

	userId := c.GetString("user_id")
	request.UserId = userId

	result, err := env.Services().Owner().Create(request)
	if err != nil {
		if err == services.ErrObjectAlreadyExists {
			c.AbortWithStatusJSON(409, sharedModels.NewErrorResponseString(c, "Owner already exists"))
			return
		}

		c.Error(err)
		c.AbortWithStatusJSON(500, sharedModels.NewErrorResponseString(c, "Internal server error"))
		return
	}

	c.JSON(201, result)
}

// GetById godoc
// @Summary Get owner profile
// @Schemes
// @Description Get owner profile
// @Tags owner
// @Produce json
// @Success 200 {object} models.Owner
// @Failure 400 {object} sharedModels.ErrorResponse
// @Failure 401 {object} sharedModels.ErrorResponse
// @Failure 404 {object} sharedModels.ErrorResponse
// @Failure 500 {object} sharedModels.ErrorResponse
// @Router /api/v1/profiles/owners [get]
func Get(env *env.Environment, c *gin.Context) {
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

	c.JSON(200, owner)
}
