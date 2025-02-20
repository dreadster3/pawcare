package mongo

import (
	"context"

	"github.com/dreadster3/pawcare/services/account/internal/owner/domain"
	"github.com/dreadster3/pawcare/shared/utils"
	"github.com/go-kit/log"
)

type ownerMiddleware func(domain.IOwnerRepository) domain.IOwnerRepository

type ownerLoggingMiddleware struct {
	logger log.Logger
	next   domain.IOwnerRepository
}

func newOwnerLoggingMiddleware(logger log.Logger) ownerMiddleware {
	return func(next domain.IOwnerRepository) domain.IOwnerRepository {
		return &ownerLoggingMiddleware{logger, next}
	}
}

func (mw *ownerLoggingMiddleware) Logger(ctx context.Context) log.Logger {
	requestId := ctx.Value(utils.RequestIdContextKey).(string)
	return log.With(mw.logger, "request_id", requestId)
}

func (mw *ownerLoggingMiddleware) FindById(ctx context.Context, id domain.OwnerId) (owner *domain.Owner, err error) {
	defer func() {
		mw.Logger(ctx).Log("method", "FindById", "id", id, "owner", owner, "err", err)
	}()
	return mw.next.FindById(ctx, id)
}

func (mw *ownerLoggingMiddleware) FindByUserId(ctx context.Context, userId domain.UserId) (owner *domain.Owner, err error) {
	defer func() {
		mw.Logger(ctx).Log("method", "FindByUserId", "userId", userId, "owner", owner, "err", err)
	}()
	return mw.next.FindByUserId(ctx, userId)
}

func (mw *ownerLoggingMiddleware) Create(ctx context.Context, owner *domain.Owner) (err error) {
	defer func() {
		mw.Logger(ctx).Log("method", "Create", "owner", owner, "err", err)
	}()

	return mw.next.Create(ctx, owner)
}

func (mw *ownerLoggingMiddleware) Update(ctx context.Context, owner *domain.Owner) (err error) {
	defer func() {
		mw.Logger(ctx).Log("method", "Update", "owner", owner, "err", err)
	}()

	return mw.next.Update(ctx, owner)
}
