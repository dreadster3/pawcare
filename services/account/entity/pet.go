package entity

type PetId string

type Pet struct {
	Id          PetId
	Name        string
	DateOfBirth string
	Species     string
	Breed       string
	Weight      float64
	Gender      EGender
}
