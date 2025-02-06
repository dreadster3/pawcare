package repository

import (
	"context"

	"github.com/dreadster3/pawcare/services/account/aggregate"
)

type IPetRepository interface {
	FindById(ctx context.Context, id aggregate.PetId) (*aggregate.Pet, error)
	FindByOwnerId(ctx context.Context, ownerId aggregate.OwnerId) ([]*aggregate.Pet, error)
	Create(ctx context.Context, pet *aggregate.Pet) error
	Update(ctx context.Context, pet *aggregate.Pet) error
}
