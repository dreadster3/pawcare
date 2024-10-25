package repository

import (
	"github.com/dreadster3/pawcare/services/profile/models"
)

type IPetRepository interface {
	FindAll() ([]models.Pet, error)
	FindById(id string) (*models.Pet, error)
	FindByOwnerId(ownerId string) ([]models.Pet, error)
	FindByIdAndOwnerId(id, ownerId string) (*models.Pet, error)
	Create(Pet models.Pet) (*models.Pet, error)
	Update(Pet models.Pet) (*models.Pet, error)
}
