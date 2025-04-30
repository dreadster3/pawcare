package events

import (
	"time"

	"github.com/dreadster3/pawcare/shared/events"
)

type PetCreated struct {
	Id          string    `json:"id"`
	UserId      string    `json:"user_id"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Species     string    `json:"species"`
	Breed       string    `json:"breed"`
	Weight      float64   `json:"weight"`
	Gender      string    `json:"gender"`
}

func (p PetCreated) EventName() string {
	return string(EventPetCreated)
}

func (p PetCreated) Key() string {
	return p.UserId
}

var _ events.IEvent = (*PetCreated)(nil)
