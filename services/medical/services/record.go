package services

import (
	"github.com/dreadster3/pawcare/services/medical/models"
	"github.com/dreadster3/pawcare/services/medical/repository"
)

type RecordService struct {
	repository repository.IRecordRepository
}

func NewRecordService(repository repository.IRecordRepository) *RecordService {
	return &RecordService{
		repository: repository,
	}
}

func (s *RecordService) Create(record models.Record) (*models.Record, error) {
	return s.repository.Create(record)
}

func (s *RecordService) FindByUserIdAndPetId(userId, petId string) ([]models.Record, error) {
	return s.repository.FindByUserIdAndPetId(userId, petId)
}

func (s *RecordService) FindByUserIdAndId(userId, id string) (*models.Record, error) {
	return s.repository.FindByUserIdAndId(userId, id)
}
