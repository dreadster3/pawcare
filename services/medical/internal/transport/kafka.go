package transport

import (
	"github.com/ThreeDotsLabs/watermill-kafka/v3/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/dreadster3/pawcare/services/medical/internal/pet/endpoint"
	"github.com/dreadster3/pawcare/services/medical/internal/pet/transport"
	"github.com/dreadster3/pawcare/shared/common"
	"github.com/dreadster3/pawcare/shared/watermill"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func NewRouter(viper viper.Viper, endpoints endpoint.Set, logger *zap.Logger) (*message.Router, error) {
	router, err := watermill.NewDefaultRouter(logger)
	if err != nil {
		return nil, err
	}

	subscriber, err := kafka.NewSubscriber(kafka.SubscriberConfig{
		Brokers:     viper.GetStringSlice(common.KafkaBrokersKey),
		Unmarshaler: kafka.DefaultMarshaler{},
	}, router.Logger())
	if err != nil {
		return nil, err
	}

	transport.RegisterKafkaRoutes(router, subscriber, endpoints)

	return router, nil
}
