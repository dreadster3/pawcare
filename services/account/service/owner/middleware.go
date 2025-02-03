package owner

import (
	"context"
	"time"

	"github.com/dreadster3/pawcare/services/account/aggregate"
	"github.com/dreadster3/pawcare/services/account/valueobjects"
	"github.com/dreadster3/pawcare/services/auth"
	"github.com/go-kit/log"
	"github.com/go-playground/validator/v10"
)

type middleware func(IOwnerService) IOwnerService

type loggingMiddleware struct {
	logger log.Logger
	next   IOwnerService
}

func newLoggingMiddleware(logger log.Logger) middleware {
	return func(next IOwnerService) IOwnerService {
		return &loggingMiddleware{logger, next}
	}
}

func (mw *loggingMiddleware) FindById(ctx context.Context, id aggregate.OwnerId) (owner *aggregate.Owner, err error) {
	defer func() {
		mw.logger.Log("method", "FindById", "id", id, "owner", owner, "err", err)
	}()

	return mw.next.FindById(ctx, id)
}

func (mw *loggingMiddleware) FindByUserId(ctx context.Context, userId auth.UserId) (*aggregate.Owner, error) {
	defer func() {
		mw.logger.Log("method", "FindByUserId", "userId", userId)
	}()

	return mw.next.FindByUserId(ctx, userId)
}

func (mw *loggingMiddleware) Save(ctx context.Context, userId auth.UserId, ownerProfile *valueobjects.OwnerProfile) error {
	defer func() {
		mw.logger.Log("method", "Save", "userId", userId, "owner", ownerProfile)
	}()

	return mw.next.Save(ctx, userId, ownerProfile)
}

type validationMiddleware struct {
	next IOwnerService
}

func newValidationMiddleware() middleware {
	return func(next IOwnerService) IOwnerService {
		return &validationMiddleware{next}
	}
}

func (mw *validationMiddleware) FindById(ctx context.Context, id aggregate.OwnerId) (*aggregate.Owner, error) {
	return mw.next.FindById(ctx, id)
}

func (mw *validationMiddleware) FindByUserId(ctx context.Context, userId auth.UserId) (*aggregate.Owner, error) {
	return mw.next.FindByUserId(ctx, userId)
}

func (mw *validationMiddleware) Save(ctx context.Context, userId auth.UserId, ownerProfile *valueobjects.OwnerProfile) error {
	if err := validator.New().Struct(ownerProfile); err != nil {
		return err
	}

	if time.Now().After(ownerProfile.DateOfBirth) || time.Now().AddDate(-150, 0, 0).Before(ownerProfile.DateOfBirth) {
		return ErrInvalidDate
	}

	return mw.next.Save(ctx, userId, ownerProfile)
}
