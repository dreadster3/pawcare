package service

import (
	"context"

	"github.com/dreadster3/pawcare/services/account/internal/owner/domain"
	"github.com/go-kit/log"
)

type IOwnerService interface {
	FindById(ctx context.Context, id domain.OwnerId) (*domain.Owner, error)
	FindByUserId(ctx context.Context, userId domain.UserId) (*domain.Owner, error)
	Create(ctx context.Context, userId domain.UserId, ownerProfile domain.OwnerProfile) (*domain.Owner, error)
	Update(ctx context.Context, owner *domain.Owner) (*domain.Owner, error)
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

func (svc *ownerService) FindById(ctx context.Context, id domain.OwnerId) (*domain.Owner, error) {
	return svc.ownerRepository.FindById(ctx, id)
}

func (svc *ownerService) FindByUserId(ctx context.Context, userId domain.UserId) (*domain.Owner, error) {
	return svc.ownerRepository.FindByUserId(ctx, userId)
}

func (svc *ownerService) Create(ctx context.Context, userId domain.UserId, ownerProfile domain.OwnerProfile) (*domain.Owner, error) {
	ownerAggregate := domain.NewOwner(userId, ownerProfile)
	if err := svc.ownerRepository.Create(ctx, ownerAggregate); err != nil {
		return nil, err
	}

	return ownerAggregate, nil
}

func (svc *ownerService) Update(ctx context.Context, owner *domain.Owner) (*domain.Owner, error) {
	if err := svc.ownerRepository.Update(ctx, owner); err != nil {
		return nil, err
	}

	return owner, nil
}
