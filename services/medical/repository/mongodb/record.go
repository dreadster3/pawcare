package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type RecordRepository struct {
	db *mongo.Database
}
