package aggregate

import "github.com/dreadster3/pawcare/services/account/entity"

type Account struct {
	User         *entity.User
	OwnerProfile *entity.Owner
	PetProfiles  []*entity.Pet
}

func (user *Account) Id() string {
	return string(user.User.Id)
}
