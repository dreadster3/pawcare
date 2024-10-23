package repository

import "github.com/dreadster3/pawcare/services/profile/models"

type IOwnerRepository interface {
	FindByUserId(userId string) (*models.Owner, error)
	Create(owner models.Owner) (*models.Owner, error)
}
