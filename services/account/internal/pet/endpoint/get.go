package endpoint

import (
	"context"
	"time"

	"github.com/dreadster3/pawcare/services/account/internal/pet/domain"
	"github.com/dreadster3/pawcare/services/account/internal/pet/service"
	"github.com/dreadster3/pawcare/shared/common"
	"github.com/dreadster3/pawcare/shared/utils"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-playground/validator/v10"
)

type GetManyResponse struct {
	Pets []GetResponse `json:"pets"`
	Err  error         `json:"-"`
}

func (r GetManyResponse) Failed() error {
	return r.Err
}

type GetByIdRequest struct {
	Id string `json:"id" validate:"required"`
}

type GetResponse struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Species     string    `json:"species"`
	Breed       string    `json:"breed"`
	Weight      float64   `json:"weight"`
	Gender      string    `json:"gender"`
	Err         error     `json:"-"`
}

func (r GetResponse) Failed() error {
	return r.Err
}

func makeGetAllEndpoint(petService service.IPetService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		pets, err := petService.GetAll(ctx)
		if err != nil {
			return GetManyResponse{Err: err}, nil
		}

		return GetManyResponse{
			Pets: utils.Map(pets, func(pet *domain.Pet) GetResponse {
				return GetResponse{
					Id:          string(pet.Id),
					Name:        pet.Profile.Name,
					Species:     pet.Profile.Species,
					Breed:       pet.Profile.Breed,
					DateOfBirth: pet.Profile.DateOfBirth,
					Weight:      pet.Profile.Weight,
					Gender:      string(pet.Profile.Gender),
				}
			}),
		}, nil
	}
}

func makeGetByIdEndpoint(petService service.IPetService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(GetByIdRequest)
		if !ok {
			return GetResponse{Err: common.ErrCastRequest}, nil
		}

		if err := validator.New().Struct(req); err != nil {
			return GetResponse{Err: err}, nil
		}

		pet, err := petService.GetById(ctx, domain.PetId(req.Id))
		if err != nil {
			return GetResponse{Err: err}, nil
		}

		return GetResponse{
			Id:          string(pet.Id),
			Name:        pet.Profile.Name,
			DateOfBirth: pet.Profile.DateOfBirth,
			Species:     pet.Profile.Species,
			Breed:       pet.Profile.Breed,
			Weight:      pet.Profile.Weight,
			Gender:      string(pet.Profile.Gender),
		}, nil
	}
}
