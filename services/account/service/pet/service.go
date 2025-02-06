package petservice

import (
	"github.com/dreadster3/pawcare/services/account/aggregate"
	"github.com/dreadster3/pawcare/services/account/repository"
	"github.com/dreadster3/pawcare/services/account/valueobjects"
	"github.com/go-kit/log"
	"golang.org/x/net/context"
)

type IPetService interface {
	FindById(ctx context.Context, petId aggregate.PetId) (*aggregate.Pet, error)
	FindByOwnerId(ctx context.Context, ownerId aggregate.OwnerId) ([]*aggregate.Pet, error)
	Create(ctx context.Context, ownerId aggregate.OwnerId, petProfile valueobjects.PetProfile) (*aggregate.Pet, error)
	Update(ctx context.Context, pet *aggregate.Pet) (*aggregate.Pet, error)
}

type petService struct {
	ownerRepository repository.IOwnerRepository
	petRepository   repository.IPetRepository
}

func NewPetService(logger log.Logger, ownerRepository repository.IOwnerRepository, petRepository repository.IPetRepository) IPetService {
	var svc IPetService
	svc = &petService{ownerRepository: ownerRepository, petRepository: petRepository}
	svc = newLoggingMiddleware(logger)(svc)

	return svc
}

func (svc *petService) FindById(ctx context.Context, petId aggregate.PetId) (*aggregate.Pet, error) {
	return svc.petRepository.FindById(ctx, petId)
}

func (svc *petService) FindByOwnerId(ctx context.Context, ownerId aggregate.OwnerId) ([]*aggregate.Pet, error) {
	return svc.petRepository.FindByOwnerId(ctx, ownerId)
}

func (svc *petService) Create(ctx context.Context, ownerId aggregate.OwnerId, petProfile valueobjects.PetProfile) (*aggregate.Pet, error) {
	petAggregate := aggregate.NewPet(ownerId, petProfile)
	if err := svc.petRepository.Create(ctx, petAggregate); err != nil {
		return nil, err
	}

	return petAggregate, nil
}

func (svc *petService) Update(ctx context.Context, pet *aggregate.Pet) (*aggregate.Pet, error) {
	if err := svc.petRepository.Update(ctx, pet); err != nil {
		return nil, err
	}

	return pet, nil
}
