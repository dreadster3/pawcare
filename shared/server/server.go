package server

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/dreadster3/pawcare/shared/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunServer(ctx context.Context, engine *gin.Engine) {
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if gin.Mode() == gin.DebugMode {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: engine,
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

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Logger.Error("Server forced shutdown", "error", err)
	}

}
