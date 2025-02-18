package pet

import (
	"fmt"
	"time"
)

type PetId string
type OwnerId string

type Pet struct {
	Id          PetId
	OwnerId     OwnerId
	Name        string
	DateOfBirth time.Time
	Species     string
	Breed       string
	Weight      float64
	Gender      string
}

func (p Pet) String() string {
	return fmt.Sprintf("Pet(id=%s, name=%s, species=%s, breed=%s, weight=%f, gender=%s)", p.Id, p.Name, p.Species, p.Breed, p.Weight, p.Gender)
}
