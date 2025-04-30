package domain

import (
	"fmt"

	"github.com/dreadster3/pawcare/services/account/internal/owner/domain"
)

type PetId string

type Pet struct {
	Id      PetId
	OwnerId domain.OwnerId
	Profile PetProfile
}

func NewPet(ownerId domain.OwnerId, profile PetProfile) (p *Pet) {
	return &Pet{
		OwnerId: ownerId,
		Profile: profile,
	}
}

func (p Pet) String() string {
	return fmt.Sprintf("Pet(id=%s,ownerId=%s,profile=%s)", p.Id, p.OwnerId, p.Profile)
}
