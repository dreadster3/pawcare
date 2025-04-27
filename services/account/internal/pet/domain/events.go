package domain

import (
	"time"

	"github.com/dreadster3/pawcare/services/account/internal/owner/domain"
	accountevents "github.com/dreadster3/pawcare/services/account/pkg/events"
	"github.com/dreadster3/pawcare/shared/events"
)

type PetCreated struct {
	Id      *PetId
	OwnerId domain.OwnerId
	Profile PetProfile

	occuredAt time.Time
}

func NewPetCreated(pet *Pet) PetCreated {
	return PetCreated{
		Id:        &pet.Id,
		OwnerId:   pet.OwnerId,
		Profile:   pet.Profile,
		occuredAt: time.Now(),
	}
}

func (petcreate PetCreated) OccurredAt() time.Time {
	return petcreate.occuredAt
}

func (p PetCreated) ToApplicationEvent() events.IEvent {
	return &accountevents.PetCreated{
		Id:          string(*p.Id),
		OwnerId:     string(p.OwnerId),
		Name:        p.Profile.Name,
		DateOfBirth: p.Profile.DateOfBirth,
		Species:     p.Profile.Species,
		Breed:       p.Profile.Breed,
		Weight:      p.Profile.Weight,
		Gender:      string(p.Profile.Gender),
	}
}

var _ events.IDomainEvent = (*PetCreated)(nil)
