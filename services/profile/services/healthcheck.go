package services

import (
	"context"
	"time"

	"github.com/dreadster3/gohealth/pkg/healthcheck"
	"github.com/dreadster3/pawcare/services/profile/db"
	"github.com/dreadster3/pawcare/services/profile/logger"
)

var HealthcheckService *healthcheck.HealthcheckService

func init() {
	HealthcheckService = healthcheck.NewHealthcheckService()

	HealthcheckService.Register("db", func(ctx context.Context, executor healthcheck.HealthcheckTaskExecutor) healthcheck.HealthcheckStatus {
		logger := logger.Logger.With("request_id", ctx.Value("request_id").(string))

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		db, disconnect, err := db.ConnectDB(ctx)
		defer disconnect(ctx)
		if err != nil {
			logger.Error("Failed to connect to database", "error", err)
			return healthcheck.HealthcheckStatusUnhealthy
		}

		if err := db.Client().Ping(ctx, nil); err != nil {
			logger.Error("Failed to ping database", "error", err)
			return healthcheck.HealthcheckStatusUnhealthy
		}

		return healthcheck.HealthcheckStatusHealthy
	})
}
