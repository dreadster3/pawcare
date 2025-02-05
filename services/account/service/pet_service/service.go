package pet

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
	Save(ctx context.Context, ownerId aggregate.OwnerId, petProfile valueobjects.PetProfile) error
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

func (svc *petService) Save(ctx context.Context, ownerId aggregate.OwnerId, petProfile valueobjects.PetProfile) error {
	petAggregate := aggregate.NewPet(ownerId, petProfile)
	return svc.petRepository.Save(ctx, petAggregate)
}
