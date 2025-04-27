package events

import (
	"context"
	"time"
)

type IEvent interface {
	EventName() string
}

type IDomainEvent interface {
	OccurredAt() time.Time
	ToApplicationEvent() IEvent
}

type IEventDispatcher interface {
	Dispatch(ctx context.Context, events []IDomainEvent) error
}

type IEventSubscriber interface {
	Subscribe(ctx context.Context, events ...IEvent)
}
