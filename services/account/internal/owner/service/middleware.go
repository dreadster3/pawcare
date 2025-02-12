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

func (mw *loggingMiddleware) Create(ctx context.Context, userId domain.UserId, ownerProfile domain.OwnerProfile) (owner *domain.Owner, err error) {
	defer func() {
		mw.logger.Log("method", "Create", "userId", userId, "ownerProfile", ownerProfile, "owner", owner, "err", err)
	}()

	return mw.next.Create(ctx, userId, ownerProfile)
}

func (mw *loggingMiddleware) Update(ctx context.Context, owner *domain.Owner) (result *domain.Owner, err error) {
	defer func() {
		mw.logger.Log("method", "Update", "owner", owner, "result", result, "err", err)
	}()
	return mw.next.Update(ctx, owner)
}

type validationMiddleware struct {
	next IOwnerService
}

func newValidationMiddleware() middleware {
	return func(next IOwnerService) IOwnerService {
		return &validationMiddleware{next}
	}
}

func (mw *validationMiddleware) FindById(ctx context.Context, id domain.OwnerId) (*domain.Owner, error) {
	return mw.next.FindById(ctx, id)
}

func (mw *validationMiddleware) FindByUserId(ctx context.Context, userId domain.UserId) (*domain.Owner, error) {
	return mw.next.FindByUserId(ctx, userId)
}

func (mw *validationMiddleware) Create(ctx context.Context, userId domain.UserId, ownerProfile domain.OwnerProfile) (*domain.Owner, error) {
	if err := validator.New().Struct(ownerProfile); err != nil {
		return nil, err
	}

	if time.Now().Before(ownerProfile.DateOfBirth) || time.Now().AddDate(-150, 0, 0).After(ownerProfile.DateOfBirth) {
		return nil, ErrInvalidDate
	}

	return mw.next.Create(ctx, userId, ownerProfile)
}

func (mw *validationMiddleware) Update(ctx context.Context, owner *domain.Owner) (*domain.Owner, error) {
	if err := validator.New().Struct(owner.Profile); err != nil {
		return nil, err
	}

	if time.Now().Before(owner.Profile.DateOfBirth) || time.Now().AddDate(-150, 0, 0).After(owner.Profile.DateOfBirth) {
		return nil, ErrInvalidDate
	}

	return mw.next.Update(ctx, owner)
}
