package service

import (
	ownerservice "github.com/dreadster3/pawcare/services/account/internal/owner/service"
	petdomain "github.com/dreadster3/pawcare/services/account/internal/pet/domain"
	"github.com/dreadster3/pawcare/shared/common"
	"github.com/dreadster3/pawcare/shared/events"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type IPetService interface {
	GetAll(ctx context.Context) ([]*petdomain.Pet, error)
	GetById(ctx context.Context, id petdomain.PetId) (*petdomain.Pet, error)
	Create(ctx context.Context, petProfile petdomain.PetProfile) (*petdomain.Pet, error)
}

type petService struct {
	petRepository   petdomain.IPetRepository
	ownerService    ownerservice.IOwnerService
	eventDispatcher events.IEventDispatcher
}

func NewPetService(petRepository petdomain.IPetRepository, ownerService ownerservice.IOwnerService, eventDispatcher events.IEventDispatcher, logger *zap.Logger) IPetService {
	var svc IPetService
	svc = &petService{petRepository: petRepository, ownerService: ownerService, eventDispatcher: eventDispatcher}
	svc = newLoggingMiddleware(logger)(svc)

	return svc
}

func (svc *petService) GetAll(ctx context.Context) ([]*petdomain.Pet, error) {
	owner, err := svc.ownerService.Get(ctx)
	if err != nil {
		return nil, err
	}

	return svc.petRepository.FindByOwnerId(ctx, owner.Id)
}

func (svc *petService) GetById(ctx context.Context, id petdomain.PetId) (*petdomain.Pet, error) {
	owner, err := svc.ownerService.Get(ctx)
	if err != nil {
		return nil, err
	}

	pet, err := svc.petRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	if pet.OwnerId != owner.Id {
		return nil, common.ErrUnauthorized
	}

	return pet, nil
}

func (svc *petService) Create(ctx context.Context, petProfile petdomain.PetProfile) (*petdomain.Pet, error) {
	owner, err := svc.ownerService.Get(ctx)
	if err != nil {
		return nil, err
	}

	pet := petdomain.NewPet(owner.Id, petProfile)
	if err := svc.petRepository.Create(ctx, pet); err != nil {
		return nil, err
	}

	if err := svc.eventDispatcher.Dispatch(ctx, pet.Events()); err != nil {
		return nil, err
	}

	return pet, nil
}
