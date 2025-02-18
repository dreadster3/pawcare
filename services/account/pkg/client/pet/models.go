package pet

import "time"

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
