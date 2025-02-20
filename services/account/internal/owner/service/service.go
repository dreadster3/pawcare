package service

import (
	"context"

	"github.com/dreadster3/pawcare/services/account/internal/owner/domain"
	"github.com/dreadster3/pawcare/shared/common"
	"github.com/go-kit/log"
)

type IOwnerService interface {
	Get(ctx context.Context) (*domain.Owner, error)
	Create(ctx context.Context, ownerProfile domain.OwnerProfile) (*domain.Owner, error)
}

type ownerService struct {
	ownerRepository domain.IOwnerRepository
}

func NewOwnerService(ownerRepository domain.IOwnerRepository, logger log.Logger) IOwnerService {
	var svc IOwnerService
	svc = &ownerService{
		ownerRepository: ownerRepository,
	}
	svc = newLoggingMiddleware(logger)(svc)
	svc = newValidationMiddleware()(svc)

	return svc
}

func (svc *ownerService) Get(ctx context.Context) (*domain.Owner, error) {
	userId, err := common.GetUserIdFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return svc.ownerRepository.FindByUserId(ctx, domain.UserId(userId))
}

func (svc *ownerService) Create(ctx context.Context, ownerProfile domain.OwnerProfile) (*domain.Owner, error) {
	userId, err := common.GetUserIdFromContext(ctx)
	if err != nil {
		return nil, err
	}

	ownerAggregate := domain.NewOwner(domain.UserId(userId), ownerProfile)
	if err := svc.ownerRepository.Create(ctx, ownerAggregate); err != nil {
		return nil, err
	}

	return ownerAggregate, nil
}
