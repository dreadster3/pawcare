package env

import (
	"github.com/dreadster3/pawcare/services/medical/services"
	sharedServices "github.com/dreadster3/pawcare/shared/services"
)

type ServiceContainer struct {
	record *services.RecordService

	auth        sharedServices.IAuthService
	profile     sharedServices.IProfileService
	healthcheck sharedServices.IHealthcheckService
}

func (c *ServiceContainer) Auth() sharedServices.IAuthService {
	return c.auth
}

func (c *ServiceContainer) Record() *services.RecordService {
	return c.record
}

func (c *ServiceContainer) Profile() sharedServices.IProfileService {
	return c.profile
}

func (c *ServiceContainer) Healthcheck() sharedServices.IHealthcheckService {
	return c.healthcheck
}
