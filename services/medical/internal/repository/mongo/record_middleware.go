package mongo

import (
	"context"

	"github.com/dreadster3/pawcare/services/medical/internal/record/domain"
	"github.com/dreadster3/pawcare/shared/utils"
	"go.uber.org/zap"
)

type recordMiddleware func(domain.IRecordRepository) domain.IRecordRepository

type loggingMiddleware struct {
	logger *zap.Logger
	next   domain.IRecordRepository
}

func newLoggingMiddleware(logger *zap.Logger) recordMiddleware {
	return func(next domain.IRecordRepository) domain.IRecordRepository {
		return &loggingMiddleware{logger, next}
	}
}

func (mw *loggingMiddleware) Logger(ctx context.Context) *zap.SugaredLogger {
	requestId := ctx.Value(utils.RequestIdContextKey).(string)
	return mw.logger.With(zap.String("request_id", requestId)).Sugar()
}

func (mw *loggingMiddleware) FindById(ctx context.Context, id domain.RecordId) (record *domain.Record, err error) {
	defer func() {
		mw.Logger(ctx).Info("method", "FindById", "id", id, "record", record, "err", err)
	}()
	return mw.next.FindById(ctx, id)
}

func (mw *loggingMiddleware) FindByPetId(ctx context.Context, petId domain.PetId) (records []*domain.Record, err error) {
	defer func() {
		mw.Logger(ctx).Info("method", "FindByPetId", "petId", petId, "records", records, "err", err)
	}()
	return mw.next.FindByPetId(ctx, petId)
}

func (mw *loggingMiddleware) Create(ctx context.Context, record *domain.Record) (err error) {
	defer func() {
		mw.Logger(ctx).Info("method", "Create", "record", record, "err", err)
	}()
	return mw.next.Create(ctx, record)
}

func (mw *loggingMiddleware) Update(ctx context.Context, record *domain.Record) (err error) {
	defer func() {
		mw.Logger(ctx).Info("method", "Update", "record", record, "err", err)
	}()
	return mw.next.Update(ctx, record)
}
