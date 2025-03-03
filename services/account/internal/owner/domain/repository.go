package domain

import (
	"context"
)

type IOwnerRepository interface {
	FindById(ctx context.Context, id OwnerId) (*Owner, error)
	FindByUserId(ctx context.Context, userId UserId) (*Owner, error)
	Create(ctx context.Context, owner *Owner) error
	Update(ctx context.Context, owner *Owner) error
}
