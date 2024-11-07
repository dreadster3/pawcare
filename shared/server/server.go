package server

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/dreadster3/pawcare/shared/config"
	"github.com/dreadster3/pawcare/shared/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupServer() *viper.Viper {
	godotenv.Load()

	viper := config.SetupConfig()

	logger.InitLogging(viper)

	return viper
}

func RunServer(ctx context.Context, viper *viper.Viper, engine *gin.Engine) error {
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if gin.Mode() == gin.DebugMode {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	port := viper.GetInt("port")

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: engine,
	}

	errorChan := make(chan error)

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logger.Logger.Error("Failed to start server", "error", err)
			errorChan <- err
		}
	}()

	logger.Logger.Info(fmt.Sprintf("Server started. Listening on port %d", port))
	logger.Logger.Info("Press Ctrl+C to stop server")

	select {
	case err := <-errorChan:
		return err
	case <-ctx.Done():
	}

	logger.Logger.Info("Shutting down server")
	logger.Logger.Info("Press Ctrl+C to force shutdown")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Logger.Error("Server forced shutdown", "error", err)
	}

	return nil
}
