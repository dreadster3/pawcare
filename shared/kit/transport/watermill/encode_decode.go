package watermill

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

type DecodeRequestFunc func(context.Context, *message.Message) (request interface{}, err error)

type EncodeRequestFunc func(context.Context, *message.Messages, interface{}) error

type EncodeResponseFunc func(context.Context, *message.Message, interface{}) error

type DecodeResponseFunc func(context.Context, *message.Message) (response interface{}, err error)
