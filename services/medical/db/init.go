package db

import (
	"context"
	"os"

	"github.com/dreadster3/pawcare/shared/db/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConnectDB(ctx context.Context) (*mongo.Database, mongodb.DbCloseFunc, error) {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "medical"
	}

	return mongodb.ConnectDB(ctx, dbName)
}
