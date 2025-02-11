package service

import (
	ownerdomain "github.com/dreadster3/pawcare/services/account/owner/domain"
	"github.com/dreadster3/pawcare/services/account/pet/domain"
	"github.com/dreadster3/pawcare/services/auth"
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

func (mw *loggingMiddleware) FindById(ctx context.Context, id domain.PetId) (pet *domain.Pet, err error) {
	defer func() {
		mw.logger.Log("method", "FindById", "id", id, "pet", pet, "err", err)
	}()
	return mw.next.FindById(ctx, id)
}

func (mw *loggingMiddleware) FindByUserId(ctx context.Context, userId auth.UserId) (pets []*domain.Pet, err error) {
	defer func() {
		mw.logger.Log("method", "FindByUserId", "userId", userId, "pets", len(pets), "err", err)
	}()
	return mw.next.FindByUserId(ctx, userId)
}

func (mw *loggingMiddleware) FindByOwnerId(ctx context.Context, ownerId ownerdomain.OwnerId) (pets []*domain.Pet, err error) {
	defer func() {
		mw.logger.Log("method", "FindByOwnerId", "ownerId", ownerId, "pets", len(pets), "err", err)
	}()
	return mw.next.FindByOwnerId(ctx, ownerId)
}

func (mw *loggingMiddleware) Create(ctx context.Context, userId auth.UserId, petProfile domain.PetProfile) (pet *domain.Pet, err error) {
	defer func() {
		mw.logger.Log("method", "Create", "userId", userId, "pet", pet, "err", err)
	}()

	return mw.next.Create(ctx, userId, petProfile)
}

func (mw *loggingMiddleware) Update(ctx context.Context, pet *domain.Pet) (p *domain.Pet, err error) {
	defer func() {
		mw.logger.Log("method", "Update", "pet", pet, "err", err)
	}()

	return mw.next.Update(ctx, pet)
}
