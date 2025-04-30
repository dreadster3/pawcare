package transport

import (
	"context"
	"fmt"

	"github.com/ThreeDotsLabs/watermill-kafka/v3/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/dreadster3/pawcare/services/account/pkg/events"
	"github.com/dreadster3/pawcare/services/medical/internal/pet/endpoint"
	"github.com/dreadster3/pawcare/shared/common"
	kitkafka "github.com/dreadster3/pawcare/shared/kafka"
)

func RegisterKafkaRoutes(router *message.Router, subscriber *kafka.Subscriber, endpoints endpoint.Set) {
	petCreateHandler := kitkafka.NewHandler(
		endpoints.CreateEndpoint,
		common.KafkaDecodeJSONMessage[endpoint.CreateRequest],
		common.KafkaEncodeResponse,
	)

	router.AddNoPublisherHandler("create_pet", fmt.Sprintf("accounts.%s", events.EventPetCreated), subscriber,
		func(msg *message.Message) error {
			return petCreateHandler(context.Background(), msg)
		})
}
