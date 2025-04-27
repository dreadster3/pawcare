package transport

import (
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v3/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/dreadster3/pawcare/services/medical/internal/pet/endpoint"
	"github.com/dreadster3/pawcare/services/medical/internal/pet/transport"
	"github.com/dreadster3/pawcare/shared/common"
	"github.com/go-kit/log"
	"github.com/spf13/viper"
)

func NewRouter(viper viper.Viper, endpoints endpoint.Set, l log.Logger) (*message.Router, error) {
	logger := watermill.NewStdLoggerWithOut(log.NewStdlibAdapter(l), false, true)
	router, err := message.NewRouter(message.RouterConfig{}, logger)
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

	subscriber, err := kafka.NewSubscriber(kafka.SubscriberConfig{
		Brokers:     viper.GetStringSlice(common.KafkaBrokersKey),
		Unmarshaler: kafka.DefaultMarshaler{},
	}, watermill.NewStdLoggerWithOut(log.NewStdlibAdapter(l), false, false))
	if err != nil {
		return nil, err
	}

	transport.RegisterKafkaRoutes(router, subscriber, endpoints)

	return router, nil
}
