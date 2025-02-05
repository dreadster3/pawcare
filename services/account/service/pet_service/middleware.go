package pet

import (
	"github.com/dreadster3/pawcare/services/account/aggregate"
	"github.com/dreadster3/pawcare/services/account/valueobjects"
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

func (mw *loggingMiddleware) FindById(ctx context.Context, id aggregate.PetId) (pet *aggregate.Pet, err error) {
	defer func() {
		mw.logger.Log("method", "FindById", "id", id, "pet", pet, "err", err)
	}()
	return mw.next.FindById(ctx, id)
}

func (mw *loggingMiddleware) FindByOwnerId(ctx context.Context, ownerId aggregate.OwnerId) (pet []*aggregate.Pet, err error) {
	defer func() {
		mw.logger.Log("method", "FindByOwnerId", "id", ownerId, "pet", pet, "err", err)
	}()
	return mw.next.FindByOwnerId(ctx, ownerId)
}

func (mw *loggingMiddleware) Save(ctx context.Context, ownerId aggregate.OwnerId, petProfile valueobjects.PetProfile) (err error) {
	defer func() {
		mw.logger.Log("method", "Save", "ownerId", ownerId, "pet", petProfile, "err", err)
	}()
	return mw.next.Save(ctx, ownerId, petProfile)
}
