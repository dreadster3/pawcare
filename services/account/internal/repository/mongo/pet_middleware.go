package mongo

import (
	"context"

	ownerdomain "github.com/dreadster3/pawcare/services/account/internal/owner/domain"
	"github.com/dreadster3/pawcare/services/account/internal/pet/domain"
	"github.com/dreadster3/pawcare/shared/utils"
	"github.com/go-kit/log"
)

type petMiddleware func(domain.IPetRepository) domain.IPetRepository

type petLoggingMiddleware struct {
	logger log.Logger
	next   domain.IPetRepository
}

func newPetLoggingMiddleware(logger log.Logger) petMiddleware {
	return func(next domain.IPetRepository) domain.IPetRepository {
		return &petLoggingMiddleware{logger, next}
	}
}

func (mw *petLoggingMiddleware) Logger(ctx context.Context) log.Logger {
	requestId := ctx.Value(utils.RequestIdContextKey).(string)
	return log.With(mw.logger, "request_id", requestId)
}

func (mw *petLoggingMiddleware) Create(ctx context.Context, pet *domain.Pet) (err error) {
	defer func() {
		mw.Logger(ctx).Log("method", "Create", "pet", pet, "err", err)
	}()

	return mw.next.Create(ctx, pet)
}

func (mw *petLoggingMiddleware) FindById(ctx context.Context, id domain.PetId) (pet *domain.Pet, err error) {
	defer func() {
		mw.Logger(ctx).Log("method", "FindById", "id", id, "pet", pet, "err", err)
	}()

	return mw.next.FindById(ctx, id)
}

func (mw *petLoggingMiddleware) Update(ctx context.Context, pet *domain.Pet) (err error) {
	defer func() {
		mw.Logger(ctx).Log("method", "Update", "pet", pet, "err", err)
	}()

	return mw.next.Update(ctx, pet)
}

func (mw *petLoggingMiddleware) FindByOwnerId(ctx context.Context, ownerId ownerdomain.OwnerId) (pets []*domain.Pet, err error) {
	defer func() {
		mw.Logger(ctx).Log("method", "FindByOwnerId", "ownerId", ownerId, "pets", pets, "err", err)
	}()

	return mw.next.FindByOwnerId(ctx, ownerId)
}
