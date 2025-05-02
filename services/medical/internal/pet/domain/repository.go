package domain

import "context"

type IPetRepository interface {
	FindById(ctx context.Context, id PetId) (*Pet, error)
	FindByUserId(ctx context.Context, userId UserId) ([]*Pet, error)
	Create(ctx context.Context, pet *Pet) error
}
