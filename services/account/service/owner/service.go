package ownerservice

import (
	"context"

	"github.com/dreadster3/pawcare/services/account/aggregate"
	"github.com/dreadster3/pawcare/services/account/repository"
	"github.com/dreadster3/pawcare/services/account/valueobjects"
	"github.com/dreadster3/pawcare/services/auth"
	"github.com/go-kit/log"
)

type IOwnerService interface {
	FindById(ctx context.Context, id aggregate.OwnerId) (*aggregate.Owner, error)
	FindByUserId(ctx context.Context, userId auth.UserId) (*aggregate.Owner, error)
	Create(ctx context.Context, userId auth.UserId, ownerProfile valueobjects.OwnerProfile) (*aggregate.Owner, error)
	Update(ctx context.Context, owner *aggregate.Owner) (*aggregate.Owner, error)
}

type ownerService struct {
	ownerRepository repository.IOwnerRepository
}

func NewOwnerService(ownerRepository repository.IOwnerRepository, logger log.Logger) IOwnerService {
	var svc IOwnerService
	svc = &ownerService{
		ownerRepository: ownerRepository,
	}
	svc = newLoggingMiddleware(logger)(svc)
	svc = newValidationMiddleware()(svc)

	return svc
}

func (svc *ownerService) FindById(ctx context.Context, id aggregate.OwnerId) (*aggregate.Owner, error) {
	return svc.ownerRepository.FindById(ctx, id)
}

func (svc *ownerService) FindByUserId(ctx context.Context, userId auth.UserId) (*aggregate.Owner, error) {
	return svc.ownerRepository.FindByUserId(ctx, userId)
}

func (svc *ownerService) Create(ctx context.Context, userId auth.UserId, ownerProfile valueobjects.OwnerProfile) (*aggregate.Owner, error) {
	ownerAggregate := aggregate.NewOwner(userId, ownerProfile)
	if err := svc.ownerRepository.Create(ctx, ownerAggregate); err != nil {
		return nil, err
	}

	return ownerAggregate, nil
}

func (svc *ownerService) Update(ctx context.Context, owner *aggregate.Owner) (*aggregate.Owner, error) {
	if err := svc.ownerRepository.Update(ctx, owner); err != nil {
		return nil, err
	}

	return owner, nil
}
