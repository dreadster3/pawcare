package env

import (
	"github.com/dreadster3/pawcare/services/medical/services"
	sharedServices "github.com/dreadster3/pawcare/shared/services"
)

type ServiceContainer struct {
	auth   *sharedServices.AuthService
	record *services.RecordService
}

func (c *ServiceContainer) Auth() sharedServices.IAuthService {
	return c.auth
}

func (c *ServiceContainer) Record() *services.RecordService {
	return c.record
}
