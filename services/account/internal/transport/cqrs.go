package transport

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v3/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/dreadster3/pawcare/shared/common"
	"github.com/dreadster3/pawcare/shared/events"
	sharedcqrs "github.com/dreadster3/pawcare/shared/watermill/cqrs"
	"github.com/dreadster3/pawcare/shared/watermill/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func NewEventBus(viper viper.Viper, logger *zap.Logger) (*cqrs.EventBus, error) {
	publisherConfig := kafka.PublisherConfig{
		Brokers: viper.GetStringSlice(common.KafkaBrokersKey),
		Marshaler: kafka.NewWithPartitioningMarshaler(func(topic string, msg *message.Message) (string, error) {
			if key, ok := msg.Context().Value("key").(string); ok {
				return key, nil
			}

			return "", nil
		}),
	}

	publisher, err := kafka.NewPublisher(publisherConfig, log.NewLogger(logger))
	if err != nil {
		return nil, err
	}

	eventBusConfig := cqrs.EventBusConfig{
		GeneratePublishTopic: func(params cqrs.GenerateEventPublishTopicParams) (string, error) {
			return "accounts." + params.EventName, nil
		},

		OnPublish: func(params cqrs.OnEventSendParams) error {
			middleware.SetCorrelationID(watermill.NewUUID(), params.Message)

			if event, ok := params.Event.(events.IEvent); ok {
				ctx := context.WithValue(params.Message.Context(), "key", event.Key())
				params.Message.SetContext(ctx)
			}

			return nil
		},
		Marshaler: cqrs.JSONMarshaler{
			GenerateName: sharedcqrs.StructNameDotLower,
		},
		Logger: log.NewLogger(logger.With(zap.String("component", "eventBus"))),
	}

	eventBus, err := cqrs.NewEventBusWithConfig(publisher, eventBusConfig)
	if err != nil {
		return nil, err
	}

	return eventBus, nil
}
