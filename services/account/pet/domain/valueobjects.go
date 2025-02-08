package domain

import "github.com/dreadster3/pawcare/services/account/owner/domain"

type PetId string

type Pet struct {
	Id      PetId
	OwnerId domain.OwnerId
	Profile PetProfile
}

func NewPet(ownerId domain.OwnerId, profile PetProfile) *Pet {
	return &Pet{
		OwnerId: ownerId,
		Profile: profile,
	}
}
