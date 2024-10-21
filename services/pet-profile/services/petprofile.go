package services

import (
	"github.com/dreadster3/pawcare/services/pet-profile/models"
	"github.com/dreadster3/pawcare/services/pet-profile/repository"
)

type PetProfileService struct {
	petProfileRepository repository.IPetProfileRepository
}

func NewPetProfileService(petProfileRepository repository.IPetProfileRepository) *PetProfileService {
	return &PetProfileService{
		petProfileRepository: petProfileRepository,
	}
}

func (s *PetProfileService) FindById(id string) (*models.PetProfile, error) {
	return s.petProfileRepository.FindById(id)
}

func (s *PetProfileService) FindAll() ([]models.PetProfile, error) {
	return s.petProfileRepository.FindAll()
}

func (s *PetProfileService) Create(petProfile models.PetProfile) (*models.PetProfile, error) {
	return s.petProfileRepository.Create(petProfile)
}

func (s *PetProfileService) Update(petProfile models.PetProfile) (*models.PetProfile, error) {
	return s.petProfileRepository.Update(petProfile)
}
