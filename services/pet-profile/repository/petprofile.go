package repository

import (
	"errors"

	"github.com/dreadster3/pawcare/services/pet-profile/models"
)

var (
	ErrProfileNotFound = errors.New("Pet profile not found")
)

type IPetProfileRepository interface {
	FindById(id string) (*models.PetProfile, error)
	FindAll() ([]models.PetProfile, error)
	Create(petProfile models.PetProfile) (*models.PetProfile, error)
	Update(petProfile models.PetProfile) (*models.PetProfile, error)
}
