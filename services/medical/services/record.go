package services

import "github.com/dreadster3/pawcare/services/medical/repository"

type RecordService struct {
	repository repository.IRecordRepository
}

func NewRecordService(repository repository.IRecordRepository) *RecordService {
	return &RecordService{
		repository: repository,
	}
}
