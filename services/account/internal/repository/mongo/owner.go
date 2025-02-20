package mongo

import (
	"context"

	"github.com/dreadster3/pawcare/services/account/internal/owner/domain"
	"github.com/dreadster3/pawcare/services/account/internal/repository"
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

func (o *owner) ToModel() *domain.Owner {
	ownerProfile := domain.NewOwnerProfile(o.Name, o.DateOfBirth.Time())
	owner := domain.NewOwner(domain.UserId(o.UserId.Hex()), ownerProfile)
	owner.Id = domain.OwnerId(o.Id.Hex())
	return owner
}

func fromOwnerModel(o domain.Owner) (*owner, error) {
	id, err := primitive.ObjectIDFromHex(string(o.Id))
	if err != nil {
		if err != primitive.ErrInvalidHex {
			return nil, err
		}
		id = primitive.NewObjectID()
	}

	userId, err := primitive.ObjectIDFromHex(string(o.UserId))
	if err != nil {
		return nil, err
	}

	return &owner{
		Id:          id,
		UserId:      userId,
		Name:        o.Profile.Name,
		DateOfBirth: primitive.NewDateTimeFromTime(o.Profile.DateOfBirth),
	}, nil
}

type ownerRepository struct {
	db *mongo.Database
}

func NewOwnerRepository(db *mongo.Database, logger log.Logger) domain.IOwnerRepository {
	var repository domain.IOwnerRepository
	repository = &ownerRepository{db}
	repository = newOwnerLoggingMiddleware(logger)(repository)

	return repository
}

func (r *ownerRepository) Collection() *mongo.Collection {
	return r.db.Collection(OwnerCollection)
}

func (r *ownerRepository) FindById(ctx context.Context, id domain.OwnerId) (*domain.Owner, error) {
	objectId, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return nil, err
	}

	var result owner
	if err := r.Collection().FindOne(ctx, bson.M{"_id": objectId}).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrNotFound
		}

		return nil, err
	}

	return result.ToModel(), nil
}

func (r *ownerRepository) FindByUserId(ctx context.Context, id domain.UserId) (*domain.Owner, error) {
	objectId, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return nil, err
	}

	var result owner
	if err := r.Collection().FindOne(ctx, bson.M{"user_id": objectId}).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrNotFound
		}

		return nil, err
	}

	return result.ToModel(), nil
}

func (r *ownerRepository) Create(ctx context.Context, owner *domain.Owner) error {
	if _, err := r.FindByUserId(ctx, owner.UserId); err == nil {
		return repository.ErrAlreadyCreated
	}

	entity, err := fromOwnerModel(*owner)
	if err != nil {
		return err
	}

	result, err := r.Collection().InsertOne(ctx, entity)
	if err != nil {
		return err
	}

	owner.Id = domain.OwnerId(result.InsertedID.(primitive.ObjectID).Hex())
	return nil
}

func (r *ownerRepository) Update(ctx context.Context, owner *domain.Owner) error {
	entity, err := fromOwnerModel(*owner)
	if err != nil {
		return err
	}

	if _, err := r.Collection().UpdateOne(ctx, bson.M{"_id": entity.Id}, bson.M{"$set": entity}); err != nil {
		return err
	}

	return nil
}
