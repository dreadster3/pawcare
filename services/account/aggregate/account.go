package aggregate

import (
	"fmt"

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

func (a Account) String() string {
	return fmt.Sprintf("Account(user=%s,owner=%s,pets=%s)", a.User, a.Owner, a.Pets)
}
