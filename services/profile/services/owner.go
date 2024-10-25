package services

import (
	"github.com/dreadster3/pawcare/services/profile/models"
	"github.com/dreadster3/pawcare/services/profile/repository"
)

type OwnerService struct {
	repository repository.IOwnerRepository
}

func NewOwnerService(repository repository.IOwnerRepository) *OwnerService {
	return &OwnerService{
		repository: repository,
	}
}

func (s *OwnerService) FindByUserId(userId string) (*models.Owner, error) {
	return s.repository.FindByUserId(userId)
}

func (s *OwnerService) Create(owner models.Owner) (*models.Owner, error) {
	if _, err := s.repository.FindByUserId(owner.UserId); err == nil {
		return nil, ErrObjectAlreadyExists
	}

	return s.repository.Create(owner)
}
