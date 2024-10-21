package mongodb

import (
	"context"

	"github.com/dreadster3/pawcare/services/pet-profile/db"
	"github.com/dreadster3/pawcare/services/pet-profile/models"
	"github.com/dreadster3/pawcare/services/pet-profile/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PetProfileRepository struct{}

func NewPetProfileRepository() *PetProfileRepository {
	return &PetProfileRepository{}
}

func (r *PetProfileRepository) FindAll() ([]models.PetProfile, error) {
	ctx := context.Background()
	database, disconnect, err := db.ConnectDB(ctx)
	defer disconnect(ctx)

	cursor, err := database.Collection("profile").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var result []models.PetProfile
	cursor.All(ctx, &result)

	return result, nil
}

func (r *PetProfileRepository) Create(profile models.PetProfile) (*models.PetProfile, error) {
	ctx := context.Background()
	database, disconnect, err := db.ConnectDB(ctx)
	defer disconnect(ctx)

	profile.Id = primitive.NewObjectID()
	result, err := database.Collection("profile").InsertOne(ctx, profile)
	if err != nil {
		return nil, err
	}

	profile.Id = result.InsertedID.(primitive.ObjectID)
	return &profile, nil
}

func (r *PetProfileRepository) FindById(id string) (*models.PetProfile, error) {
	ctx := context.Background()
	database, disconnect, err := db.ConnectDB(ctx)
	defer disconnect(ctx)

	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var result models.PetProfile
	err = database.Collection("profile").FindOne(ctx, bson.M{"_id": idObj}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrProfileNotFound
		}

		return nil, err
	}

	return &result, nil
}

func (r *PetProfileRepository) Update(profile models.PetProfile) (*models.PetProfile, error) {
	ctx := context.Background()
	database, disconnect, err := db.ConnectDB(ctx)
	defer disconnect(ctx)

	_, err = database.Collection("profile").UpdateOne(ctx, bson.M{"_id": profile.Id}, bson.M{"$set": profile})
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *PetProfileRepository) Delete(id string) error {
	ctx := context.Background()
	database, disconnect, err := db.ConnectDB(ctx)
	defer disconnect(ctx)

	_, err = database.Collection("profile").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
