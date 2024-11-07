package mongodb

import (
	"context"

	"github.com/dreadster3/pawcare/services/medical/models"
	"github.com/dreadster3/pawcare/services/medical/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecordRepository struct {
	db *mongo.Database
}

func NewRecordRepository(db *mongo.Database) *RecordRepository {
	return &RecordRepository{db: db}
}

func (r *RecordRepository) Create(record models.Record) (*models.Record, error) {
	ctx := context.Background()

	record.Id = primitive.NewObjectID()
	result, err := r.db.Collection(RECORD_COLLECTION).InsertOne(ctx, record)
	if err != nil {
		return nil, err
	}

	record.Id = result.InsertedID.(primitive.ObjectID)
	return &record, nil
}

func (r *RecordRepository) FindByUserIdAndPetId(userId string, petId string) ([]models.Record, error) {
	ctx := context.Background()

	cursor, err := r.db.Collection(RECORD_COLLECTION).Find(ctx, bson.M{"user_id": userId, "pet_id": petId})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrRecordNotFound
		}

		return nil, err
	}

	var results []models.Record
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return []models.Record{}, nil
	}

	return results, nil
}

func (r *RecordRepository) FindByUserIdAndId(userId, id string) (*models.Record, error) {
	ctx := context.Background()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, repository.ErrInvalidId
	}

	var record models.Record
	err = r.db.Collection(RECORD_COLLECTION).FindOne(ctx, bson.M{"user_id": userId, "_id": objectId}).Decode(&record)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrRecordNotFound
		}
		return nil, err
	}

	return &record, nil
}
