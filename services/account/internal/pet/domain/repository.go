package domain

import (
	"context"

	"github.com/dreadster3/pawcare/services/account/internal/owner/domain"
)

type IPetRepository interface {
	FindById(ctx context.Context, id PetId) (*Pet, error)
	FindByOwnerId(ctx context.Context, ownerId domain.OwnerId) ([]*Pet, error)
	Create(ctx context.Context, pet *Pet) error
	Update(ctx context.Context, pet *Pet) error
}
