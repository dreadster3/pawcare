package repository

import "github.com/dreadster3/pawcare/services/medical/models"

type IRecordRepository interface {
	Create(record models.Record) (*models.Record, error)
	FindByUserIdAndPetId(userId string, petId string) ([]models.Record, error)
	FindByUserIdAndId(userId string, id string) (*models.Record, error)
}
