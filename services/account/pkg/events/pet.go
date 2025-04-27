package events

import (
	"time"

	"github.com/dreadster3/pawcare/shared/events"
)

type PetCreated struct {
	Id          string
	OwnerId     string
	Name        string
	DateOfBirth time.Time
	Species     string
	Breed       string
	Weight      float64
	Gender      string
}

func (p PetCreated) EventName() string {
	return string(PetCreatedEvent)
}

var _ events.IEvent = (*PetCreated)(nil)
