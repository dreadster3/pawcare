package endpoint

import (
	"context"
	"time"

	"github.com/dreadster3/pawcare/services/medical/internal/record/domain"
	"github.com/dreadster3/pawcare/services/medical/internal/record/service"
	"github.com/dreadster3/pawcare/shared/common"
	"github.com/dreadster3/pawcare/shared/utils"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-playground/validator/v10"
)

type GetByPetIdRequest struct {
	PetId string `json:"id" validate:"required"`
}

type GetResponse struct {
	Id          string    `json:"id"`
	Type        string    `json:"type"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

func makeGetByPetIdEndpoint(recordService service.IRecordService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(GetByPetIdRequest)
		if !ok {
			return nil, common.ErrCastRequest
		}

		if err := validator.New().Struct(req); err != nil {
			return nil, err
		}

		petId := domain.PetId(req.PetId)
		records, err := recordService.FindByPetId(ctx, petId)
		if err != nil {
			return nil, err
		}

		return utils.Map(records, func(r *domain.Record) GetResponse {
			return GetResponse{
				Id:          string(r.Id),
				Type:        string(r.RecordInfo.Type),
				Date:        r.RecordInfo.Date,
				Description: r.RecordInfo.Description,
			}
		}), nil
	}
}
