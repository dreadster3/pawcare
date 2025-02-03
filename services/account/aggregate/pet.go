package aggregate

import "github.com/dreadster3/pawcare/services/account/valueobjects"

type PetId string

type Pet struct {
	Id      PetId
	OwnerId OwnerId
	Profile valueobjects.PetProfile
}
