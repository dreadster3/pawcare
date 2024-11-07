package profiles

import (
	"github.com/dreadster3/pawcare/shared/config"
	"github.com/dreadster3/pawcare/shared/models/profile"
	"github.com/dreadster3/pawcare/shared/services"
	"github.com/dreadster3/pawcare/shared/services/http"
	"github.com/dreadster3/pawcare/shared/services/http/profiles/client"
	"github.com/dreadster3/pawcare/shared/services/http/profiles/client/pet"
	"github.com/gin-gonic/gin"
)

type ProfileService struct {
	petProfileService *PetProfileService
}

func NewProfileService(serviceConfig *config.ServiceConfig, ctx *gin.Context) *ProfileService {
	httpContext := http.NewServiceContext(serviceConfig.Address(), ctx)

	return &ProfileService{NewPetProfileService(httpContext)}
}

func (s *ProfileService) Pet() services.IPetProfileService {
	return s.petProfileService
}

func (s *ProfileService) Owner() services.IOwnerProfileService {
	return nil
}

func (s *ProfileService) Healthcheck() services.IHealthcheckService {
	return nil
}

type PetProfileService struct {
	ctx    *http.ServiceContext
	client *client.ProfileService
}

func NewPetProfileService(ctx *http.ServiceContext) *PetProfileService {
	transportConfig := client.DefaultTransportConfig().WithHost(ctx.Host)
	client := client.NewHTTPClientWithConfig(nil, transportConfig)
	return &PetProfileService{ctx, client}
}

func (s *PetProfileService) FindById(id string) (*profile.ModelsPet, error) {
	token := s.ctx.BearerToken()
	response, err := s.client.Pet.GetAPIV1ProfilesPetsID(pet.NewGetAPIV1ProfilesPetsIDParams().WithID(id), http.BearerToken(token))
	if err != nil {
		if _, ok := err.(*pet.GetAPIV1ProfilesPetsIDNotFound); ok {
			return nil, services.ErrNotFound
		}

		return nil, err
	}

	return response.Payload, nil
}
