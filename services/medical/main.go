package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/dreadster3/pawcare/services/medical/api"
	"github.com/dreadster3/pawcare/services/medical/db"
	"github.com/dreadster3/pawcare/services/medical/env"
	"github.com/dreadster3/pawcare/shared/logger"
	"github.com/dreadster3/pawcare/shared/server"
	"github.com/joho/godotenv"
)

func _main() error {
	godotenv.Load()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	db, disconnect, err := db.ConnectDB(ctx)
	defer disconnect(ctx)
	if err != nil {
		return err
	}

	environment := env.InitEnvironment(db)

	engine := server.NewDefaultEngine()

	api.RegisterRoutes(environment, &engine.RouterGroup)

	server.RunServer(ctx, engine)

	return nil
}

func main() {
	if err := _main(); err != nil {
		logger.Logger.Error("Fatal error", "error", err)
		panic(err)
	}
}
