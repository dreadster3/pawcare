package domain

import (
	"context"

	"github.com/dreadster3/pawcare/services/auth"
)

type IOwnerRepository interface {
	FindById(ctx context.Context, id OwnerId) (*Owner, error)
	FindByUserId(ctx context.Context, userId auth.UserId) (*Owner, error)
	Create(ctx context.Context, owner *Owner) error
	Update(ctx context.Context, owner *Owner) error
}
