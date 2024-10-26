package mongodb

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/dreadster3/pawcare/shared/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbCloseFunc func(context.Context) error

func clientOptions() *options.ClientOptions {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "27017"
	}

	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)
	clientOptions := options.Client().ApplyURI(connectionString)

	return clientOptions
}

func createClient(ctx context.Context, clientOptions *options.ClientOptions) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func ConnectDB(ctx context.Context, databaseName string) (*mongo.Database, DbCloseFunc, error) {
	clientOptions := clientOptions()

	logger.Logger.Debug("Connecting to database")
	client, err := createClient(ctx, clientOptions)
	if err != nil {
		return nil, func(ctx context.Context) error { return nil }, err
	}

	disconnect := func(ctx context.Context) error {
		logger.Logger.Debug("Disconnecting from database")
		return client.Disconnect(ctx)
	}

	return client.Database(databaseName), disconnect, nil
}
