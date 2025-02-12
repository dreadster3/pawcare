package middleware

import (
	"context"

	"github.com/dreadster3/pawcare/services/account/internal/owner/domain"
	"github.com/go-kit/log"
)

type ownerMiddleware func(domain.IOwnerRepository) domain.IOwnerRepository

func OwnerWarpMiddleware(repository domain.IOwnerRepository, logger log.Logger) domain.IOwnerRepository {
	var result domain.IOwnerRepository = repository
	result = newLoggingMiddleware(logger)(result)
	return result
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.IOwnerRepository
}

func newLoggingMiddleware(logger log.Logger) ownerMiddleware {
	return func(next domain.IOwnerRepository) domain.IOwnerRepository {
		return &loggingMiddleware{logger, next}
	}
}

func (mw *loggingMiddleware) FindById(ctx context.Context, id domain.OwnerId) (owner *domain.Owner, err error) {
	defer func() {
		mw.logger.Log("method", "FindById", "id", id, "owner", owner, "err", err)
	}()
	return mw.next.FindById(ctx, id)
}

func (mw *loggingMiddleware) FindByUserId(ctx context.Context, userId domain.UserId) (owner *domain.Owner, err error) {
	defer func() {
		mw.logger.Log("method", "FindByUserId", "userId", userId, "owner", owner, "err", err)
	}()
	return mw.next.FindByUserId(ctx, userId)
}

func (mw *loggingMiddleware) Create(ctx context.Context, owner *domain.Owner) (err error) {
	defer func() {
		mw.logger.Log("method", "Create", "owner", owner, "err", err)
	}()

	return mw.next.Create(ctx, owner)
}

func (mw *loggingMiddleware) Update(ctx context.Context, owner *domain.Owner) (err error) {
	defer func() {
		mw.logger.Log("method", "Update", "owner", owner, "err", err)
	}()

	return mw.next.Update(ctx, owner)
}
