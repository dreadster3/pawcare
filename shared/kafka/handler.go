package kafka

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-kit/kit/endpoint"
)

type (
	DecodeRequestFunc func(context.Context, *message.Message) (interface{}, error)
	EncodeResponse    func(context.Context, interface{}) error
)

func NewHandler(endpoint endpoint.Endpoint, decode DecodeRequestFunc, encode EncodeResponse) func(context.Context, *message.Message) error {
	return func(ctx context.Context, msg *message.Message) error {
		req, err := decode(ctx, msg)
		if err != nil {
			return err
		}

		response, err := endpoint(ctx, req)
		if err != nil {
			return err
		}

		return encode(ctx, response)
	}
}
