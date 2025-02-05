package aggregate

import (
	"github.com/dreadster3/pawcare/services/account/valueobjects"
	"github.com/dreadster3/pawcare/services/auth"
)

type OwnerId string

type Owner struct {
	Id      OwnerId
	UserId  auth.UserId
	Profile valueobjects.OwnerProfile
}

func NewOwner(userId auth.UserId, profile valueobjects.OwnerProfile) *Owner {
	return &Owner{
		UserId:  userId,
		Profile: profile,
	}
}
