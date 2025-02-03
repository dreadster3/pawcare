package owner

import (
	"context"

	"github.com/dreadster3/pawcare/services/account/aggregate"
	"github.com/dreadster3/pawcare/services/account/repository"
	"github.com/dreadster3/pawcare/services/account/valueobjects"
	"github.com/dreadster3/pawcare/services/auth"
)

type IOwnerService interface {
	FindById(ctx context.Context, id aggregate.OwnerId) (*aggregate.Owner, error)
	FindByUserId(ctx context.Context, userId auth.UserId) (*aggregate.Owner, error)
	Save(ctx context.Context, userId auth.UserId, ownerProfile *valueobjects.OwnerProfile) error
}

type ownerService struct {
	ownerRepository repository.IOwnerRepository
	userService     auth.IUserService
}

func NewOwnerService(ownerRepository repository.IOwnerRepository, userService auth.IUserService) IOwnerService {
	return &ownerService{
		ownerRepository: ownerRepository,
		userService:     userService,
	}
}

func (svc *ownerService) FindById(ctx context.Context, id aggregate.OwnerId) (*aggregate.Owner, error) {
	return svc.ownerRepository.FindById(ctx, id)
}

func (svc *ownerService) FindByUserId(ctx context.Context, userId auth.UserId) (*aggregate.Owner, error) {
	return svc.ownerRepository.FindByUserId(ctx, userId)
}

func (svc *ownerService) Save(ctx context.Context, userId auth.UserId, ownerProfile *valueobjects.OwnerProfile) error {
	ownerAggregate := aggregate.NewOwner(userId, ownerProfile)
	return svc.ownerRepository.Save(ctx, ownerAggregate)
}
