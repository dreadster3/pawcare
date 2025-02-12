package service

import (
	"context"

	"github.com/dreadster3/pawcare/services/medical/internal/record/domain"
	"github.com/go-kit/log"
)

type IRecordService interface {
	FindById(ctx context.Context, id domain.RecordId) (*domain.Record, error)
	FindByPetId(ctx context.Context, petId domain.PetId) ([]*domain.Record, error)
	Create(ctx context.Context, petId domain.PetId, record domain.RecordInfo) (*domain.Record, error)
	Update(ctx context.Context, record *domain.Record) error
}

type recordService struct {
	recordRepository domain.IRecordRepository
}

func NewRecordService(recordRepository domain.IRecordRepository, logger log.Logger) IRecordService {
	var svc IRecordService
	svc = &recordService{recordRepository: recordRepository}
	svc = newLoggingMiddleware(logger)(svc)

	return svc
}

func (svc *recordService) FindById(ctx context.Context, id domain.RecordId) (*domain.Record, error) {
	return svc.recordRepository.FindById(ctx, id)
}

func (svc *recordService) FindByPetId(ctx context.Context, petId domain.PetId) ([]*domain.Record, error) {
	return svc.recordRepository.FindByPetId(ctx, petId)
}

func (svc *recordService) Create(ctx context.Context, petId domain.PetId, recordInfo domain.RecordInfo) (*domain.Record, error) {
	recordAggregate := domain.NewRecord(petId, recordInfo)
	if err := svc.recordRepository.Create(ctx, recordAggregate); err != nil {
		return nil, err
	}

	return recordAggregate, nil
}

func (svc *recordService) Update(ctx context.Context, record *domain.Record) error {
	return svc.recordRepository.Update(ctx, record)
}
