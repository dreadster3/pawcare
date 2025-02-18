package endpoint

import (
	"context"
	"time"

	"github.com/dreadster3/pawcare/services/account/pkg/client/pet"
	"github.com/dreadster3/pawcare/services/medical/internal/record/domain"
	"github.com/dreadster3/pawcare/services/medical/internal/record/service"
	"github.com/dreadster3/pawcare/shared/common"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-playground/validator/v10"
)

type CreateRequest struct {
	PetId       string    `json:"pet_id" validate:"required"`
	Type        string    `json:"type" validate:"required,oneof=vaccination treatment deworming surgery checkup other"`
	Description string    `json:"description" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
}

type CreateResponse struct {
	Id string `json:"id"`
}

func makeCreateEndpoint(recordService service.IRecordService, petService pet.IPetService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(CreateRequest)
		if !ok {
			return nil, common.ErrCastRequest
		}

		if err := validator.New().Struct(req); err != nil {
			return nil, err
		}

		if _, err := petService.GetById(ctx, pet.PetId(req.PetId)); err != nil {
			return nil, err
		}

		petId := domain.PetId(req.PetId)
		recordInfo := domain.NewRecordInfo(domain.RecordType(req.Type), req.Description, req.Date)
		record, err := recordService.Create(ctx, petId, recordInfo)
		if err != nil {
			return nil, err
		}

		return CreateResponse{string(record.Id)}, nil
	}
}
