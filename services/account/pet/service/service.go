package service

import (
	ownerdomain "github.com/dreadster3/pawcare/services/account/owner/domain"
	ownerservice "github.com/dreadster3/pawcare/services/account/owner/service"
	"github.com/dreadster3/pawcare/services/account/pet/domain"
	"github.com/dreadster3/pawcare/services/auth"
	"github.com/go-kit/log"
	"golang.org/x/net/context"
)

type IPetService interface {
	FindById(ctx context.Context, petId domain.PetId) (*domain.Pet, error)
	FindByUserId(ctx context.Context, userId auth.UserId) ([]*domain.Pet, error)
	FindByOwnerId(ctx context.Context, ownerId ownerdomain.OwnerId) ([]*domain.Pet, error)
	Create(ctx context.Context, userId auth.UserId, petProfile domain.PetProfile) (*domain.Pet, error)
	Update(ctx context.Context, pet *domain.Pet) (*domain.Pet, error)
}

type petService struct {
	petRepository domain.IPetRepository
	ownerService  ownerservice.IOwnerService
}

func NewPetService(petRepository domain.IPetRepository, ownerService ownerservice.IOwnerService, logger log.Logger) IPetService {
	var svc IPetService
	svc = &petService{petRepository: petRepository, ownerService: ownerService}
	svc = newLoggingMiddleware(logger)(svc)

	return svc
}

func (svc *petService) FindById(ctx context.Context, petId domain.PetId) (*domain.Pet, error) {
	return svc.petRepository.FindById(ctx, petId)
}

func (svc *petService) FindByUserId(ctx context.Context, userId auth.UserId) ([]*domain.Pet, error) {
	owner, err := svc.ownerService.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	return svc.FindByOwnerId(ctx, owner.Id)
}

func (svc *petService) FindByOwnerId(ctx context.Context, ownerId ownerdomain.OwnerId) ([]*domain.Pet, error) {
	return svc.petRepository.FindByOwnerId(ctx, ownerId)
}

func (svc *petService) Create(ctx context.Context, userId auth.UserId, petProfile domain.PetProfile) (*domain.Pet, error) {
	owner, err := svc.ownerService.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	petdomain := domain.NewPet(owner.Id, petProfile)
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
