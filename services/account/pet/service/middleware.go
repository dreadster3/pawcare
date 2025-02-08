package service

import (
	ownerdomain "github.com/dreadster3/pawcare/services/account/owner/domain"
	"github.com/dreadster3/pawcare/services/account/pet/domain"
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

func (mw *loggingMiddleware) FindByOwnerId(ctx context.Context, ownerId ownerdomain.OwnerId) (pet []*domain.Pet, err error) {
	defer func() {
		mw.logger.Log("method", "FindByOwnerId", "id", ownerId, "pet", pet, "err", err)
	}()
	return mw.next.FindByOwnerId(ctx, ownerId)
}

func (mw *loggingMiddleware) Create(ctx context.Context, ownerId ownerdomain.OwnerId, petProfile domain.PetProfile) (pet *domain.Pet, err error) {
	defer func() {
		mw.logger.Log("method", "Create", "ownerId", ownerId, "pet", pet, "err", err)
	}()

	return mw.next.Create(ctx, ownerId, petProfile)
}

func (mw *loggingMiddleware) Update(ctx context.Context, pet *domain.Pet) (p *domain.Pet, err error) {
	defer func() {
		mw.logger.Log("method", "Update", "pet", pet, "err", err)
	}()

	return mw.next.Update(ctx, pet)
}
