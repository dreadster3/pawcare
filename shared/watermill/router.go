package watermill

import (
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/dreadster3/pawcare/shared/watermill/log"
	"go.uber.org/zap"
)

func NewDefaultRouter(logger *zap.Logger) (*message.Router, error) {
	router, err := message.NewRouter(message.RouterConfig{}, log.NewLogger(logger))
	if err != nil {
		return nil, err
	}

	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(
		middleware.CorrelationID,
		middleware.Retry{
			MaxRetries:      3,
			InitialInterval: 50 * time.Millisecond,
		}.Middleware,
		middleware.Recoverer,
	)

	return router, nil
}
