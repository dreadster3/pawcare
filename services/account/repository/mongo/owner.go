package mongo

import (
	"context"

	"github.com/dreadster3/pawcare/services/account/aggregate"
	"github.com/dreadster3/pawcare/services/account/repository"
	"github.com/dreadster3/pawcare/services/account/valueobjects"
	"github.com/dreadster3/pawcare/services/auth"
	"github.com/go-kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	OwnerCollection = "owners"
)

type owner struct {
	Id          primitive.ObjectID `bson:"_id"`
	UserId      primitive.ObjectID `bson:"user_id"`
	Name        string             `bson:"name"`
	DateOfBirth primitive.DateTime `bson:"date_of_birth"`
}

func (o *owner) ToOwner() *aggregate.Owner {
	return &aggregate.Owner{
		Id: aggregate.OwnerId(o.Id.String()),
		Profile: valueobjects.OwnerProfile{
			Name:        o.Name,
			DateOfBirth: o.DateOfBirth.Time(),
		},
	}
}

func FromOwner(userId auth.UserId, o aggregate.Owner) (*owner, error) {
	id, err := primitive.ObjectIDFromHex(string(o.Id))
	if err != nil {
		id = primitive.NewObjectID()
	}

	userObjectId, err := primitive.ObjectIDFromHex(string(userId))
	if err != nil {
		return nil, err
	}

	return &owner{
		Id:          id,
		UserId:      userObjectId,
		Name:        o.Profile.Name,
		DateOfBirth: primitive.NewDateTimeFromTime(o.Profile.DateOfBirth),
	}, nil
}

type ownerRepository struct {
	db *mongo.Database
}

func NewOwnerRepository(logger log.Logger, db *mongo.Database) repository.IOwnerRepository {
	var repo repository.IOwnerRepository
	repo = &ownerRepository{db}
	repo = newLoggingMiddleware(logger)(repo)

	return repo
}

func (r *ownerRepository) FindById(ctx context.Context, id aggregate.OwnerId) (*aggregate.Owner, error) {
	objectId, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return nil, err
	}

	var result owner
	if err := r.db.Collection(OwnerCollection).FindOne(ctx, bson.M{"_id": objectId}).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return result.ToOwner(), nil
}

func (r *ownerRepository) FindByUserId(ctx context.Context, id auth.UserId) (*aggregate.Owner, error) {
	objectId, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return nil, err
	}

	var result owner
	if err := r.db.Collection(OwnerCollection).FindOne(ctx, bson.M{"user_id": objectId}).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return result.ToOwner(), nil
}

func (r *ownerRepository) Create(ctx context.Context, userId auth.UserId, owner *aggregate.Owner) error {
	if _, err := r.FindByUserId(ctx, userId); err == nil {
		return repository.ErrAlreadyCreated
	}

	dbEntity, err := FromOwner(userId, *owner)
	if err != nil {
		return err
	}

	result, err := r.db.Collection(OwnerCollection).InsertOne(ctx, dbEntity)
	if err != nil {
		return err
	}

	owner.Id = aggregate.OwnerId(result.InsertedID.(primitive.ObjectID).Hex())

	return nil
}
