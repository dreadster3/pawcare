package service

import (
	"context"
	"time"

	"github.com/dreadster3/pawcare/services/account/internal/owner/domain"
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

func (mw *loggingMiddleware) Get(ctx context.Context) (owner *domain.Owner, err error) {
	defer func() {
		mw.logger.Log("method", "Get", "owner", owner, "err", err)
	}()

	return mw.next.Get(ctx)
}

func (mw *loggingMiddleware) Create(ctx context.Context, ownerProfile domain.OwnerProfile) (owner *domain.Owner, err error) {
	defer func() {
		mw.logger.Log("method", "Create", "ownerProfile", ownerProfile, "owner", owner, "err", err)
	}()

	return mw.next.Create(ctx, ownerProfile)
}

type validationMiddleware struct {
	next IOwnerService
}

func newValidationMiddleware() middleware {
	return func(next IOwnerService) IOwnerService {
		return &validationMiddleware{next}
	}
}

func (mw *validationMiddleware) Get(ctx context.Context) (*domain.Owner, error) {
	return mw.next.Get(ctx)
}

func (mw *validationMiddleware) Create(ctx context.Context, ownerProfile domain.OwnerProfile) (*domain.Owner, error) {
	if err := validator.New().Struct(ownerProfile); err != nil {
		return nil, err
	}

	if time.Now().Before(ownerProfile.DateOfBirth) || time.Now().AddDate(-150, 0, 0).After(ownerProfile.DateOfBirth) {
		return nil, ErrInvalidDate
	}

	return mw.next.Create(ctx, ownerProfile)
}
