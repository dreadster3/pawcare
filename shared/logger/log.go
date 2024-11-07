package logger

import (
	"log/slog"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var Logger *slog.Logger

func InitLogging(viper *viper.Viper) {
	level := slog.LevelInfo
	switch strings.ToLower(viper.GetString("log_level")) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn", "warning":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		if gin.Mode() == gin.DebugMode {
			level = slog.LevelDebug
		} else {
			level = slog.LevelInfo
		}
	}

	Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	}))
}

func WithRequestId(c *gin.Context) *slog.Logger {
	requestId := c.GetString("request_id")
	return Logger.With("request_id", requestId)
}
