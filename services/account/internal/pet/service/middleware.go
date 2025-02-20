package service

import (
	"github.com/dreadster3/pawcare/services/account/internal/pet/domain"
	"github.com/dreadster3/pawcare/shared/utils"
	"github.com/go-kit/log"
	"golang.org/x/net/context"
)

type middleware func(IPetService) IPetService

type loggingMiddleware struct {
	logger log.Logger
	next   IPetService
}

func newLoggingMiddleware(logger log.Logger) middleware {
	return func(next IPetService) IPetService {
		return &loggingMiddleware{logger, next}
	}
}

func (mw *loggingMiddleware) Logger(ctx context.Context) log.Logger {
	requestId := ctx.Value(utils.RequestIdContextKey).(string)
	return log.With(mw.logger, "request_id", requestId)
}

func (mw *loggingMiddleware) GetById(ctx context.Context, id domain.PetId) (pet *domain.Pet, err error) {
	defer func() {
		mw.Logger(ctx).Log("method", "GetById", "id", id, "pet", pet, "err", err)
	}()
	return mw.next.GetById(ctx, id)
}

func (mw *loggingMiddleware) GetAll(ctx context.Context) (pets []*domain.Pet, err error) {
	defer func() {
		mw.Logger(ctx).Log("method", "GetAll", "pets", len(pets), "err", err)
	}()
	return mw.next.GetAll(ctx)
}

func (mw *loggingMiddleware) Create(ctx context.Context, petProfile domain.PetProfile) (pet *domain.Pet, err error) {
	defer func() {
		mw.Logger(ctx).Log("method", "Create", "pet", pet, "err", err)
	}()

	return mw.next.Create(ctx, petProfile)
}
