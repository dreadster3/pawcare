//go:generate swag init -g main.go -o ./docs -parseDependency -parseInternal
package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/dreadster3/pawcare/services/profile/api"
	_ "github.com/dreadster3/pawcare/services/profile/docs"
	"github.com/dreadster3/pawcare/shared/db/mongodb"
	"github.com/dreadster3/pawcare/shared/logger"
	"github.com/dreadster3/pawcare/shared/server"
)

// @title           Profile Service
// @version         1.0
// @description     Service for managing pet and owner profiles

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
func _main() error {
	viper := server.SetupServer()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	db, disconnect, err := mongodb.ConnectDB(ctx, viper)
	defer disconnect(ctx)
	if err != nil {
		return err
	}

	engine := server.NewDefaultEngine()

	api.RegisterRoutes(viper, db, &engine.RouterGroup)

	return server.RunServer(ctx, viper, engine)
}

func main() {
	if err := _main(); err != nil {
		logger.Logger.Error("Fatal error", "error", err)
		panic(err)
	}
}
