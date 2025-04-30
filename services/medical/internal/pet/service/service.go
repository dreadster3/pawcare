package service

import (
	"context"
	"errors"

	"github.com/dreadster3/pawcare/services/medical/internal/pet/domain"
	"go.uber.org/zap"
)

type IPetService interface {
	Create(ctx context.Context, id domain.PetId, userId domain.UserId) (*domain.Pet, error)
}

func NewPetService(logger *zap.Logger) IPetService {
	return &petService{
		logger: logger,
	}
}

type petService struct {
	logger *zap.Logger
}

func (s *petService) Create(ctx context.Context, id domain.PetId, userId domain.UserId) (*domain.Pet, error) {
	s.logger.Info("HEREHEREHERE")
	return nil, errors.New("not implemented")
}
