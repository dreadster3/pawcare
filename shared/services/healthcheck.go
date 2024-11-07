package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dreadster3/gohealth/pkg/healthcheck"
	"github.com/dreadster3/pawcare/shared/config"
	"github.com/dreadster3/pawcare/shared/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type IHealthcheckService interface {
	Run(c *gin.Context) healthcheck.HealthcheckReport
}

type healthcheckService struct {
	*healthcheck.HealthcheckService

	db *mongo.Database
}

func NewHealthcheckService(db *mongo.Database) *healthcheckService {
	service := &healthcheckService{
		HealthcheckService: healthcheck.NewHealthcheckService(),
		db:                 db,
	}

	service.RegisterDatabase()

	return service
}

func (s *healthcheckService) Register(name string, task healthcheck.HealthcheckFn) {
	s.HealthcheckService.Register(name, task)
}

func (s *healthcheckService) RegisterDatabase() {
	s.Register("db", func(ctx context.Context, executor healthcheck.HealthcheckTaskExecutor) healthcheck.HealthcheckStatus {
		logger := logger.Logger.With("request_id", ctx.Value("request_id").(string))

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		if err := s.db.Client().Ping(ctx, nil); err != nil {
			logger.Error("Failed to ping database", "error", err)
			return healthcheck.HealthcheckStatusUnhealthy
		}

		return healthcheck.HealthcheckStatusHealthy
	})
}

func (s *healthcheckService) Run(c *gin.Context) healthcheck.HealthcheckReport {
	return s.HealthcheckService.Run(c)
}

func (s *healthcheckService) RegisterServiceCheck(name string, serviceConfig *config.ServiceConfig) {
	s.Register(name, func(ctx context.Context, executor healthcheck.HealthcheckTaskExecutor) healthcheck.HealthcheckStatus {
		response, err := http.Get(fmt.Sprintf("%s/api/healthcheck", serviceConfig.Address()))
		if err != nil {
			logger.Logger.Error("Failed to make healthcheck request", "error", err)
			return healthcheck.HealthcheckStatusDegraded
		}

		defer response.Body.Close()

		jsonString, err := io.ReadAll(response.Body)
		if err != nil {
			logger.Logger.Error("Failed to read response body", "error", err)
			return healthcheck.HealthcheckStatusDegraded
		}

		var obj healthcheck.JSONHealthcheckReport

		if err := json.Unmarshal(jsonString, &obj); err != nil {
			logger.Logger.Error("Failed to unmarshal response body", "error", err)
			return healthcheck.HealthcheckStatusDegraded
		}

		if obj.Status == healthcheck.HealthcheckStatusUnhealthy {
			return healthcheck.HealthcheckStatusDegraded
		}

		return obj.Status
	})
}
