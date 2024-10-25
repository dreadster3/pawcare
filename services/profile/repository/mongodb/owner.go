package mongodb

import (
	"context"

	"github.com/dreadster3/pawcare/services/profile/models"
	"github.com/dreadster3/pawcare/services/profile/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OwnerRepository struct {
	db *mongo.Database
}

func NewOwnerRepository(db *mongo.Database) *OwnerRepository {
	return &OwnerRepository{db}
}

func (r *OwnerRepository) FindByUserId(userId string) (*models.Owner, error) {
	ctx := context.Background()

	var result models.Owner
	err := r.db.Collection(OWNER_COLLECTION).FindOne(ctx, bson.M{"user_id": userId}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrNotFound
		}

		return nil, err
	}

	return &result, nil
}

func (r *OwnerRepository) Create(owner models.Owner) (*models.Owner, error) {
	ctx := context.Background()

	owner.Id = primitive.NewObjectID()
	result, err := r.db.Collection(OWNER_COLLECTION).InsertOne(ctx, owner)
	if err != nil {
		return nil, err
	}

	owner.Id = result.InsertedID.(primitive.ObjectID)
	return &owner, nil
}
