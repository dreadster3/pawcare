package transport

import (
	"github.com/ThreeDotsLabs/watermill-kafka/v3/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/dreadster3/pawcare/services/medical/internal/pet/endpoint"
	"github.com/dreadster3/pawcare/services/medical/internal/pet/transport"
	"github.com/dreadster3/pawcare/shared/common"
	"github.com/dreadster3/pawcare/shared/watermill/router"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func NewRouter(viper viper.Viper, endpoints endpoint.Set, logger *zap.Logger) (*message.Router, error) {
	r, err := router.NewDefaultRouter(logger)
	if err != nil {
		return nil, err
	}

	kafkaSubscriber, err := kafka.NewSubscriber(kafka.SubscriberConfig{
		Brokers:     viper.GetStringSlice(common.KafkaBrokersKey),
		Unmarshaler: kafka.DefaultMarshaler{},
	}, r.Logger())
	if err != nil {
		return nil, err
	}

	kafkaPublisher, err := kafka.NewPublisher(kafka.PublisherConfig{
		Brokers:   viper.GetStringSlice(common.KafkaBrokersKey),
		Marshaler: router.NewDefaultPartitionKeyMarshaler(),
	}, r.Logger())
	if err != nil {
		return nil, err
	}

	if err := transport.RegisterKafkaRoutes(r, kafkaSubscriber, kafkaPublisher, endpoints); err != nil {
		return nil, err
	}

	return r, nil
}
