package repository

import (
	"context"

	"github.com/dreadster3/pawcare/services/account/aggregate"
	"github.com/dreadster3/pawcare/services/auth"
	"go.mongodb.org/mongo-driver/mongo"
)

type IOwnerRepository interface {
	FindById(ctx context.Context, id aggregate.OwnerId) (*aggregate.Owner, error)
	FindByUserId(ctx context.Context, userId auth.UserId) (*aggregate.Owner, error)
	Save(ctx context.Context, owner *aggregate.Owner) error
}

type ownerRepository struct {
	db *mongo.Database
}

func (r *ownerRepository) FindById(ctx context.Context, id aggregate.OwnerId) (*aggregate.Owner, error) {
	panic("implement me")
}

func (r *ownerRepository) FindByUserId(ctx context.Context, id auth.UserId) (*aggregate.Owner, error) {
	panic("implement me")
}

func (r *ownerRepository) Save(ctx context.Context, owner *aggregate.Owner) error {
	panic("implement me")
}
