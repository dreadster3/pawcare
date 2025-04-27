package events

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaEventDispatcher struct {
	writer *kafka.Writer
	topic  string
}

func NewKafkaEventDispatcher(brokers []string, topic string) *KafkaEventDispatcher {
	return &KafkaEventDispatcher{
		topic: topic,
		writer: &kafka.Writer{
			Addr:                   kafka.TCP(brokers...),
			Topic:                  topic,
			Balancer:               &kafka.LeastBytes{},
			AllowAutoTopicCreation: true,
			MaxAttempts:            3,
			BatchTimeout:           10 * time.Millisecond,
			Async:                  true,
		},
	}
}

func (dispatcher *KafkaEventDispatcher) writeMessages(ctx context.Context, messages ...kafka.Message) error {
	for range max(dispatcher.writer.MaxAttempts, 1) {
		if err := dispatcher.writer.WriteMessages(ctx, messages...); err != nil {
			if errors.Is(err, kafka.UnknownTopicOrPartition) {
				// Retry when topic does not exist to allow creation
				time.Sleep(50 * time.Millisecond)
				continue
			}

			return err
		}

		return nil
	}

	return errors.New("Unreachable")
}

func (dispatcher *KafkaEventDispatcher) Dispatch(ctx context.Context, events []IDomainEvent) error {
	msgs := make([]kafka.Message, 0, len(events))
	for _, event := range events {
		applicationEvent := event.ToApplicationEvent()
		payload, err := json.Marshal(applicationEvent) // or Protobuf / Avro
		if err != nil {
			return err
		}

		msgs = append(msgs, kafka.Message{
			Key:   []byte(applicationEvent.EventName()),
			Value: payload,
			Time:  event.OccurredAt(),
			Headers: []kafka.Header{
				{Key: "name", Value: []byte(applicationEvent.EventName())},
				{Key: "content-type", Value: []byte("application/json")},
			},
		})
	}

	// Retry when topic does not exist to allow creation
	return dispatcher.writeMessages(ctx, msgs...)
}

func (dispatcher *KafkaEventDispatcher) Close() error {
	return dispatcher.writer.Close()
}

var _ IEventDispatcher = (*KafkaEventDispatcher)(nil)

type KafkaEventConsumer struct {
	reader *kafka.Reader
}

func (c *KafkaEventConsumer) Subscribe(ctx context.Context, events ...IEvent) {
}

var _ IEventSubscriber = (*KafkaEventConsumer)(nil)
