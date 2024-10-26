package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/dreadster3/pawcare/services/profile/api"
	"github.com/dreadster3/pawcare/services/profile/db"
	docs "github.com/dreadster3/pawcare/services/profile/docs"
	"github.com/dreadster3/pawcare/services/profile/env"
	"github.com/dreadster3/pawcare/shared/logger"
	"github.com/dreadster3/pawcare/shared/middleware"
	"github.com/dreadster3/pawcare/shared/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func recovery(c *gin.Context, err any) {
	c.AbortWithStatusJSON(500, models.NewErrorResponseString(c, "Internal server error"))
}

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

	router := gin.New()
	router.Use(gin.Logger(), gin.CustomRecovery(recovery), middleware.RequestIdMiddleware)

	api.RegisterRoutes(environment, &router.RouterGroup)

	if gin.Mode() == gin.DebugMode {
		docs.SwaggerInfo.Title = "Pet Profile Service"
		docs.SwaggerInfo.Description = "This is the Pet Profile Service"
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logger.Logger.Error("Failed to start server", "error", err)
		}
	}()

	logger.Logger.Info("Server started. Listening on port 8080")
	logger.Logger.Info("Press Ctrl+C to stop server")

	<-ctx.Done()

	logger.Logger.Info("Shutting down server")
	logger.Logger.Info("Press Ctrl+C to force shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Logger.Error("Server forced shutdown", "error", err)
	}

	return nil
}

func main() {
	if err := _main(); err != nil {
		logger.Logger.Error("Fatal error", "error", err)
		panic(err)
	}
}
