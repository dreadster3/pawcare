package endpoint

import (
	"context"
	"time"

	"github.com/dreadster3/pawcare/services/medical/internal/record/domain"
	"github.com/dreadster3/pawcare/services/medical/internal/record/service"
	"github.com/dreadster3/pawcare/shared/common"
	"github.com/go-kit/kit/endpoint"
)

type CreateRequest struct {
	PetId       string            `json:"pet_id" validate:"required"`
	RecordType  domain.RecordType `json:"record_type" validate:"required"`
	Description string            `json:"description" validate:"required"`
	Date        time.Time         `json:"date" validate:"required"`
}

type CreateResponse struct {
	Id string `json:"id"`
}

func makeCreateEndpoint(recordService service.IRecordService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(CreateRequest)
		if !ok {
			return nil, common.ErrCastRequest
		}

		petId := domain.PetId(req.PetId)
		recordInfo := domain.NewRecordInfo(req.RecordType, req.Description, req.Date)
		record, err := recordService.Create(ctx, petId, recordInfo)
		if err != nil {
			return nil, err
		}

		return CreateResponse{string(record.Id)}, nil
	}
}
