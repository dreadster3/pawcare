package profile

import (
	"errors"

	"github.com/dreadster3/pawcare/services/pet-profile/env"
	"github.com/dreadster3/pawcare/services/pet-profile/models"
	"github.com/dreadster3/pawcare/services/pet-profile/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create godoc
// @Summary Create a new pet profile
// @Schemes
// @Description Creates a new pet profile
// @Tags pet
// @Accept json
// @Param pet_profile body models.PetProfile true "Pet Profile"
// @Produce json
// @Success 201 {object} models.PetProfile
// @Failure 400 {object} string "Reason for failure"
// @Failure 500 {object} string "Reason for failure"
// @Router /api/v1/profiles [post]
func Create(env *env.Environment, c *gin.Context) {
	var request models.PetProfile
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(400, gin.H{"error": err})
		return
	}

	result, err := env.Services.PetProfile.Create(request)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(201, result)
}

// GetAll godoc
// @Summary Get all pet profiles
// @Schemes
// @Description Get all pet profiles
// @Tags pet
// @Produce json
// @Success 200 {array} models.PetProfile
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/profiles [get]
func GetAll(env *env.Environment, c *gin.Context) {
	result, err := env.Services.PetProfile.FindAll()
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, result)
}

// GetByID godoc
// @Summary Get pet profile by ID
// @Schemes
// @Description Get pet profile by ID
// @Tags pet
// @Param id path string true "Pet Profile ID"
// @Produce json
// @Success 200 {object} models.PetProfile
// @Failure 400 {object} string "Reason for failure"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/profiles/{id} [get]
func GetByID(env *env.Environment, c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithError(400, errors.New("ID is required"))
		return
	}

	result, err := env.Services.PetProfile.FindById(id)
	if err != nil {
		if err == primitive.ErrInvalidHex {
			c.Error(err)
			c.AbortWithStatusJSON(400, gin.H{"error": "Invalid ID"})
			return
		}

		if err == repository.ErrProfileNotFound {
			c.Error(err)
			c.AbortWithStatusJSON(404, gin.H{"error": "Profile not found"})
			return
		}

		c.Error(err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, result)
}
