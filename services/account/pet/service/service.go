package service

import (
	ownerdomain "github.com/dreadster3/pawcare/services/account/owner/domain"
	"github.com/dreadster3/pawcare/services/account/pet/domain"
	"github.com/go-kit/log"
	"golang.org/x/net/context"
)

type IPetService interface {
	FindById(ctx context.Context, petId domain.PetId) (*domain.Pet, error)
	FindByOwnerId(ctx context.Context, ownerId ownerdomain.OwnerId) ([]*domain.Pet, error)
	Create(ctx context.Context, ownerId ownerdomain.OwnerId, petProfile domain.PetProfile) (*domain.Pet, error)
	Update(ctx context.Context, pet *domain.Pet) (*domain.Pet, error)
}

type petService struct {
	petRepository domain.IPetRepository
}

func NewPetService(petRepository domain.IPetRepository, logger log.Logger) IPetService {
	var svc IPetService
	svc = &petService{petRepository: petRepository}
	svc = newLoggingMiddleware(logger)(svc)

	return svc
}

func (svc *petService) FindById(ctx context.Context, petId domain.PetId) (*domain.Pet, error) {
	return svc.petRepository.FindById(ctx, petId)
}

func (svc *petService) FindByOwnerId(ctx context.Context, ownerId ownerdomain.OwnerId) ([]*domain.Pet, error) {
	return svc.petRepository.FindByOwnerId(ctx, ownerId)
}

func (svc *petService) Create(ctx context.Context, ownerId ownerdomain.OwnerId, petProfile domain.PetProfile) (*domain.Pet, error) {
	petdomain := domain.NewPet(ownerId, petProfile)
	if err := svc.petRepository.Create(ctx, petdomain); err != nil {
		return nil, err
	}

	return petdomain, nil
}

func (svc *petService) Update(ctx context.Context, pet *domain.Pet) (*domain.Pet, error) {
	if err := svc.petRepository.Update(ctx, pet); err != nil {
		return nil, err
	}

	return pet, nil
}
