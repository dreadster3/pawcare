package services

import (
	"github.com/dreadster3/pawcare/services/profile/models"
	"github.com/dreadster3/pawcare/services/profile/repository"
)

type PetService struct {
	petRepository repository.IPetRepository
}

func NewPetService(petRepository repository.IPetRepository) *PetService {
	return &PetService{
		petRepository: petRepository,
	}
}

func (s *PetService) FindById(id string) (*models.Pet, error) {
	return s.petRepository.FindById(id)
}

func (s *PetService) FindAll() ([]models.Pet, error) {
	return s.petRepository.FindAll()
}

func (s *PetService) Create(pet models.Pet) (*models.Pet, error) {
	return s.petRepository.Create(pet)
}

func (s *PetService) Update(pet models.Pet) (*models.Pet, error) {
	return s.petRepository.Update(pet)
}

func (s *PetService) FindByOwnerId(ownerId string) ([]models.Pet, error) {
	return s.petRepository.FindByOwnerId(ownerId)
}

func (s *PetService) FindByIdAndOwnerId(id, ownerId string) (*models.Pet, error) {
	return s.petRepository.FindByIdAndOwnerId(id, ownerId)
}
