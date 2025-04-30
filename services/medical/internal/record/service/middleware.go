package service

import (
	"context"

	"github.com/dreadster3/pawcare/services/medical/internal/record/domain"
	"github.com/dreadster3/pawcare/shared/utils"
	"go.uber.org/zap"
)

type middleware func(IRecordService) IRecordService

type loggingMiddleware struct {
	logger *zap.Logger
	next   IRecordService
}

func newLoggingMiddleware(logger *zap.Logger) middleware {
	return func(next IRecordService) IRecordService {
		return &loggingMiddleware{logger, next}
	}
}

func (mw *loggingMiddleware) Logger(ctx context.Context) *zap.SugaredLogger {
	requestId := ctx.Value(utils.RequestIdContextKey).(string)
	return mw.logger.With(zap.String("request_id", requestId)).Sugar()
}

func (mw *loggingMiddleware) GetById(ctx context.Context, id domain.RecordId) (record *domain.Record, err error) {
	defer func() {
		mw.Logger(ctx).Info("method", "GetById", "id", id, "record", record, "err", err)
	}()

	return mw.next.GetById(ctx, id)
}

func (mw *loggingMiddleware) GetByPetId(ctx context.Context, petId domain.PetId) (records []*domain.Record, err error) {
	defer func() {
		mw.Logger(ctx).Info("method", "GetByPetId", "petId", petId, "records", len(records), "err", err)
	}()

	return mw.next.GetByPetId(ctx, petId)
}

func (mw *loggingMiddleware) Create(ctx context.Context, petId domain.PetId, recordInfo domain.RecordInfo) (record *domain.Record, err error) {
	defer func() {
		mw.Logger(ctx).Info("method", "Create", "petId", petId, "recordInfo", recordInfo, "err", err)
	}()

	return mw.next.Create(ctx, petId, recordInfo)
}
