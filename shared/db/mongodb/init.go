package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/dreadster3/pawcare/shared/logger"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbCloseFunc func(context.Context) error

func clientOptions(viper *viper.Viper) *options.ClientOptions {
	user := viper.GetString("db_user")
	password := viper.GetString("db_password")
	host := viper.GetString("db_host")
	port := viper.GetInt("db_port")

	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%d", user, password, host, port)
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

func ConnectDB(ctx context.Context, viper *viper.Viper) (*mongo.Database, DbCloseFunc, error) {
	clientOptions := clientOptions(viper)

	logger.Logger.Debug("Connecting to database")
	client, err := createClient(ctx, clientOptions)
	if err != nil {
		return nil, func(ctx context.Context) error { return nil }, err
	}

	disconnect := func(ctx context.Context) error {
		logger.Logger.Debug("Disconnecting from database")
		return client.Disconnect(ctx)
	}

	return client.Database(viper.GetString("db_name")), disconnect, nil
}
