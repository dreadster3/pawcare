package entity

import "fmt"

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

func (p Pet) String() string {
	return fmt.Sprintf("Pet(id=%s,name=%s,dob=%s,species=%s,breed=%s,weight=%f,gender=%s)", p.Id, p.Name, p.DateOfBirth, p.Species, p.Breed, p.Weight, p.Gender)
}
