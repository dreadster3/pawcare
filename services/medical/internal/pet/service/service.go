package service

import (
	"context"
	"errors"

	"github.com/dreadster3/pawcare/services/medical/internal/pet/domain"
	"github.com/go-kit/log"
)

type IPetService interface {
	Create(ctx context.Context, id domain.PetId, userId domain.UserId) (*domain.Pet, error)
}

func NewPetService(logger log.Logger) IPetService {
	return &petService{
		logger: logger,
	}
}

type petService struct {
	logger log.Logger
}

func (s *petService) Create(ctx context.Context, id domain.PetId, userId domain.UserId) (*domain.Pet, error) {
	s.logger.Log("HEREHEREHERE")
	return nil, errors.New("not implemented")
}
