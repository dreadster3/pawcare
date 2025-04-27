package domain

import (
	"fmt"

	"github.com/dreadster3/pawcare/services/account/internal/owner/domain"
	"github.com/dreadster3/pawcare/shared/events"
)

type PetId string

type Pet struct {
	Id      PetId
	OwnerId domain.OwnerId
	Profile PetProfile
	events  []events.IDomainEvent
}

func NewPet(ownerId domain.OwnerId, profile PetProfile) (p *Pet) {
	defer func() { p.pushEvent(NewPetCreated(p)) }()
	return &Pet{
		OwnerId: ownerId,
		Profile: profile,
	}
}

func (p *Pet) pushEvent(e events.IDomainEvent) {
	p.events = append(p.events, e)
}

func (p Pet) Events() []events.IDomainEvent {
	defer func() { p.events = []events.IDomainEvent{} }()
	return p.events
}

func (p Pet) String() string {
	return fmt.Sprintf("Pet(id=%s,ownerId=%s,profile=%s)", p.Id, p.OwnerId, p.Profile)
}
