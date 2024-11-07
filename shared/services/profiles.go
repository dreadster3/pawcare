package services

import "github.com/dreadster3/pawcare/shared/models/profile"

type IPetProfileService interface {
	FindById(id string) (*profile.ModelsPet, error)
}

type IOwnerProfileService interface{}

type IProfileService interface {
	Pet() IPetProfileService
	Owner() IOwnerProfileService
	Healthcheck() IHealthcheckService
}
