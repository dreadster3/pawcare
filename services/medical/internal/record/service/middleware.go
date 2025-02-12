package service

import (
	"context"

	"github.com/dreadster3/pawcare/services/medical/internal/record/domain"
	"github.com/go-kit/log"
)

type middleware func(IRecordService) IRecordService

type loggingMiddleware struct {
	logger log.Logger
	next   IRecordService
}

func newLoggingMiddleware(logger log.Logger) middleware {
	return func(next IRecordService) IRecordService {
		return &loggingMiddleware{logger, next}
	}
}

func (mw *loggingMiddleware) FindById(ctx context.Context, id domain.RecordId) (record *domain.Record, err error) {
	defer func() {
		mw.logger.Log("method", "FindById", "id", id, "record", record, "err", err)
	}()

	return mw.next.FindById(ctx, id)
}

func (mw *loggingMiddleware) FindByPetId(ctx context.Context, petId domain.PetId) (records []*domain.Record, err error) {
	defer func() {
		mw.logger.Log("method", "FindByPetId", "petId", petId, "records", len(records), "err", err)
	}()

	return mw.next.FindByPetId(ctx, petId)
}

func (mw *loggingMiddleware) Create(ctx context.Context, petId domain.PetId, recordInfo domain.RecordInfo) (record *domain.Record, err error) {
	defer func() {
		mw.logger.Log("method", "Create", "petId", petId, "recordInfo", recordInfo, "err", err)
	}()

	return mw.next.Create(ctx, petId, recordInfo)
}

func (mw *loggingMiddleware) Update(ctx context.Context, record *domain.Record) (err error) {
	defer func() {
		mw.logger.Log("method", "Update", "record", record, "err", err)
	}()

	return mw.next.Update(ctx, record)
}
