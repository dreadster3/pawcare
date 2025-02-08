package domain

import (
	"github.com/dreadster3/pawcare/services/auth"
)

type OwnerId string

type Owner struct {
	Id      OwnerId
	UserId  auth.UserId
	Profile OwnerProfile
}

func NewOwner(userId auth.UserId, profile OwnerProfile) *Owner {
	return &Owner{
		UserId:  userId,
		Profile: profile,
	}
}
