package mongodb

import (
	"context"

	"github.com/dreadster3/pawcare/services/profile/db"
	"github.com/dreadster3/pawcare/services/profile/models"
	"github.com/dreadster3/pawcare/services/profile/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PetRepository struct{}

func NewPetRepository() *PetRepository {
	return &PetRepository{}
}

func (r *PetRepository) FindAll() ([]models.Pet, error) {
	ctx := context.Background()
	database, disconnect, err := db.ConnectDB(ctx)
	defer disconnect(ctx)
	if err != nil {
		return nil, err
	}

	cursor, err := database.Collection(PET_COLLECTION).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var result []models.Pet
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return []models.Pet{}, nil
	}

	return result, nil
}

func (r *PetRepository) Create(profile models.Pet) (*models.Pet, error) {
	ctx := context.Background()
	database, disconnect, err := db.ConnectDB(ctx)
	defer disconnect(ctx)
	if err != nil {
		return nil, err
	}

	profile.Id = primitive.NewObjectID()
	result, err := database.Collection(PET_COLLECTION).InsertOne(ctx, profile)
	if err != nil {
		return nil, err
	}

	profile.Id = result.InsertedID.(primitive.ObjectID)
	return &profile, nil
}

func (r *PetRepository) FindById(id string) (*models.Pet, error) {
	ctx := context.Background()
	database, disconnect, err := db.ConnectDB(ctx)
	defer disconnect(ctx)
	if err != nil {
		return nil, err
	}

	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		if err == primitive.ErrInvalidHex {
			return nil, repository.ErrInvalidId
		}
		return nil, err
	}

	var result models.Pet
	err = database.Collection(PET_COLLECTION).FindOne(ctx, bson.M{"_id": idObj}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrNotFound
		}

		return nil, err
	}

	return &result, nil
}

func (r *PetRepository) Update(profile models.Pet) (*models.Pet, error) {
	ctx := context.Background()
	database, disconnect, err := db.ConnectDB(ctx)
	defer disconnect(ctx)
	if err != nil {
		return nil, err
	}

	_, err = database.Collection(PET_COLLECTION).UpdateOne(ctx, bson.M{"_id": profile.Id}, bson.M{"$set": profile})
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *PetRepository) Delete(id string) error {
	ctx := context.Background()
	database, disconnect, err := db.ConnectDB(ctx)
	defer disconnect(ctx)
	if err != nil {
		return err
	}

	_, err = database.Collection(PET_COLLECTION).DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}

func (r *PetRepository) FindByOwnerId(ownerId string) ([]models.Pet, error) {
	ctx := context.Background()
	database, disconnect, err := db.ConnectDB(ctx)
	defer disconnect(ctx)
	if err != nil {
		return nil, err
	}

	cursor, err := database.Collection(PET_COLLECTION).Find(ctx, bson.M{"owner_id": ownerId})
	if err != nil {
		return nil, err
	}

	var result []models.Pet
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return []models.Pet{}, nil
	}

	return result, nil
}

func (r *PetRepository) FindByIdAndOwnerId(id, ownerId string) (*models.Pet, error) {
	ctx := context.Background()
	database, disconnect, err := db.ConnectDB(ctx)
	defer disconnect(ctx)
	if err != nil {
		return nil, err
	}

	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		if err == primitive.ErrInvalidHex {
			return nil, repository.ErrInvalidId
		}
		return nil, err
	}

	var result models.Pet
	err = database.Collection(PET_COLLECTION).FindOne(ctx, bson.M{"_id": idObj, "owner_id": ownerId}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return &result, nil
}
