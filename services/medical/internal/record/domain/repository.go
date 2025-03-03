package domain

import "context"

type IRecordRepository interface {
	FindById(ctx context.Context, id RecordId) (*Record, error)
	FindByPetId(ctx context.Context, petId PetId) ([]*Record, error)
	Create(ctx context.Context, record *Record) error
	Update(ctx context.Context, record *Record) error
}
