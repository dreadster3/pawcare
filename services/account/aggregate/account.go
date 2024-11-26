package aggregate

import (
	"github.com/dreadster3/pawcare/services/account/entity"
	"github.com/dreadster3/pawcare/services/auth"
)

type Account struct {
	User  auth.User
	Owner *entity.Owner
	Pets  []*entity.Pet
}

func (user *Account) Id() string {
	return string(user.User.Id)
}
