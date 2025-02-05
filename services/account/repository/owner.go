package repository

import (
	"context"

	"github.com/dreadster3/pawcare/services/account/aggregate"
	"github.com/dreadster3/pawcare/services/auth"
)

type IOwnerRepository interface {
	FindById(ctx context.Context, id aggregate.OwnerId) (*aggregate.Owner, error)
	FindByUserId(ctx context.Context, userId auth.UserId) (*aggregate.Owner, error)
	Create(ctx context.Context, owner *aggregate.Owner) error
	Update(ctx context.Context, owner *aggregate.Owner) error
}
