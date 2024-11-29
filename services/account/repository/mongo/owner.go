package mongo

import (
	"context"

	"github.com/dreadster3/pawcare/services/account/entity"
	"github.com/dreadster3/pawcare/services/account/repository"
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

func (o *owner) ToOwner() entity.Owner {
	return entity.Owner{
		Id:          entity.OwnerId(o.Id.String()),
		Name:        o.Name,
		DateOfBirth: o.DateOfBirth.Time(),
	}
}

func FromOwner(userId auth.UserId, o entity.Owner) (*owner, error) {
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
		Name:        o.Name,
		DateOfBirth: primitive.NewDateTimeFromTime(o.DateOfBirth),
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

func (r *ownerRepository) FindById(ctx context.Context, id string) (entity.Owner, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.Owner{}, err
	}

	var result owner
	if err := r.db.Collection(OwnerCollection).FindOne(ctx, bson.M{"_id": objectId}).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return entity.Owner{}, repository.ErrNotFound
		}
		return entity.Owner{}, err
	}

	return result.ToOwner(), nil
}

func (r *ownerRepository) FindByUserId(ctx context.Context, id auth.UserId) (entity.Owner, error) {
	objectId, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return entity.Owner{}, err
	}

	var result owner
	if err := r.db.Collection(OwnerCollection).FindOne(ctx, bson.M{"user_id": objectId}).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return entity.Owner{}, repository.ErrNotFound
		}
		return entity.Owner{}, err
	}

	return result.ToOwner(), nil
}

func (r *ownerRepository) Create(ctx context.Context, userId auth.UserId, owner *entity.Owner) error {
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

	owner.Id = entity.OwnerId(result.InsertedID.(primitive.ObjectID).Hex())

	return nil
}
