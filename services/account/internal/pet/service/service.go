package service

import (
	ownerdomain "github.com/dreadster3/pawcare/services/account/internal/owner/domain"
	ownerservice "github.com/dreadster3/pawcare/services/account/internal/owner/service"
	petdomain "github.com/dreadster3/pawcare/services/account/internal/pet/domain"
	"github.com/go-kit/log"
	"golang.org/x/net/context"
)

type IPetService interface {
	FindById(ctx context.Context, petId petdomain.PetId) (*petdomain.Pet, error)
	FindByUserId(ctx context.Context, userId ownerdomain.UserId) ([]*petdomain.Pet, error)
	FindByOwnerId(ctx context.Context, ownerId ownerdomain.OwnerId) ([]*petdomain.Pet, error)
	Create(ctx context.Context, userId ownerdomain.UserId, petProfile petdomain.PetProfile) (*petdomain.Pet, error)
	Update(ctx context.Context, pet *petdomain.Pet) (*petdomain.Pet, error)
}

type petService struct {
	petRepository petdomain.IPetRepository
	ownerService  ownerservice.IOwnerService
}

func NewPetService(petRepository petdomain.IPetRepository, ownerService ownerservice.IOwnerService, logger log.Logger) IPetService {
	var svc IPetService
	svc = &petService{petRepository: petRepository, ownerService: ownerService}
	svc = newLoggingMiddleware(logger)(svc)

	return svc
}

func (svc *petService) FindById(ctx context.Context, petId petdomain.PetId) (*petdomain.Pet, error) {
	return svc.petRepository.FindById(ctx, petId)
}

func (svc *petService) FindByUserId(ctx context.Context, userId ownerdomain.UserId) ([]*petdomain.Pet, error) {
	owner, err := svc.ownerService.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	return svc.FindByOwnerId(ctx, owner.Id)
}

func (svc *petService) FindByOwnerId(ctx context.Context, ownerId ownerdomain.OwnerId) ([]*petdomain.Pet, error) {
	return svc.petRepository.FindByOwnerId(ctx, ownerId)
}

func (svc *petService) Create(ctx context.Context, userId ownerdomain.UserId, petProfile petdomain.PetProfile) (*petdomain.Pet, error) {
	owner, err := svc.ownerService.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	petpetdomain := petdomain.NewPet(owner.Id, petProfile)
	if err := svc.petRepository.Create(ctx, petpetdomain); err != nil {
		return nil, err
	}

	return petpetdomain, nil
}

func (svc *petService) Update(ctx context.Context, pet *petdomain.Pet) (*petdomain.Pet, error) {
	if err := svc.petRepository.Update(ctx, pet); err != nil {
		return nil, err
	}

	return pet, nil
}
