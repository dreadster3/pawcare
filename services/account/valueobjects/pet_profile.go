package valueobjects

import "fmt"

type PetProfile struct {
	Name        string
	DateOfBirth string
	Species     string
	Breed       string
	Weight      float64
	Gender      EGender
}

func (p PetProfile) String() string {
	return fmt.Sprintf("PetProfile(name=%s,dob=%s,species=%s,breed=%s,weight=%f,gender=%s)", p.Name, p.DateOfBirth, p.Species, p.Breed, p.Weight, p.Gender)
}
