package aggregate

import "github.com/dreadster3/pawcare/services/account/valueobjects"

type PetId string

type Pet struct {
	Id      PetId
	OwnerId OwnerId
	Profile valueobjects.PetProfile
}

func NewPet(ownerId OwnerId, profile valueobjects.PetProfile) *Pet {
	return &Pet{
		OwnerId: ownerId,
		Profile: profile,
	}
}
