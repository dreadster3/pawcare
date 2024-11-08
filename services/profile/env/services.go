package env

import (
	"github.com/dreadster3/pawcare/services/profile/services"
	sharedServices "github.com/dreadster3/pawcare/shared/services"
)

type ServiceContainer struct {
	pet   *services.PetService
	owner *services.OwnerService

	auth        sharedServices.IAuthService
	healthcheck sharedServices.IHealthcheckService
}

func (c *ServiceContainer) Auth() sharedServices.IAuthService {
	return c.auth
}

func (c *ServiceContainer) Pet() *services.PetService {
	return c.pet
}

func (c *ServiceContainer) Owner() *services.OwnerService {
	return c.owner
}

func (c *ServiceContainer) Healthcheck() sharedServices.IHealthcheckService {
	return c.healthcheck
}
