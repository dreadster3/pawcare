package repository

import (
	"context"

	"github.com/dreadster3/pawcare/services/account/entity"
	"github.com/dreadster3/pawcare/services/auth"
)

type IOwnerRepository interface {
	FindById(ctx context.Context, id string) (entity.Owner, error)
	FindByUserId(ctx context.Context, id auth.UserId) (entity.Owner, error)
	Create(ctx context.Context, userId auth.UserId, owner *entity.Owner) error
}
