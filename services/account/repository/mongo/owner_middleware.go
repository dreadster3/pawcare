package mongo

import (
	"context"

	"github.com/dreadster3/pawcare/services/account/aggregate"
	"github.com/dreadster3/pawcare/services/account/repository"
	"github.com/dreadster3/pawcare/services/auth"
	"github.com/go-kit/log"
)

type ownerMiddleware func(repository.IOwnerRepository) repository.IOwnerRepository

type loggingMiddleware struct {
	logger log.Logger
	next   repository.IOwnerRepository
}

func newLoggingMiddleware(logger log.Logger) ownerMiddleware {
	return func(next repository.IOwnerRepository) repository.IOwnerRepository {
		return &loggingMiddleware{logger, next}
	}
}

func (mw *loggingMiddleware) FindById(ctx context.Context, id aggregate.OwnerId) (owner *aggregate.Owner, err error) {
	defer func() {
		mw.logger.Log("method", "FindById", "id", id, "owner", owner, "err", err)
	}()

	return mw.next.FindById(ctx, id)
}

func (mw *loggingMiddleware) FindByUserId(ctx context.Context, id auth.UserId) (owner *aggregate.Owner, err error) {
	defer func() {
		mw.logger.Log("method", "FindByUserId", "id", id, "owner", owner, "err", err)
	}()
	return mw.next.FindByUserId(ctx, id)
}

func (mw *loggingMiddleware) Create(ctx context.Context, userId auth.UserId, owner *aggregate.Owner) (err error) {
	defer func() {
		mw.logger.Log("method", "Create", "owner", owner, "err", err)
	}()

	return mw.next.Create(ctx, userId, owner)
}
