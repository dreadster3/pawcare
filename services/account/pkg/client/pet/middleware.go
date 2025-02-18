package pet

import (
	"context"

	"github.com/go-kit/log"
)

type middleware func(IPetService) IPetService

type loggingMiddleware struct {
	logger log.Logger
	next   IPetService
}

func newLoggingMiddleware(logger log.Logger) middleware {
	return func(next IPetService) IPetService {
		return &loggingMiddleware{logger, next}
	}
}

func (mw *loggingMiddleware) GetById(ctx context.Context, id PetId) (pet Pet, err error) {
	defer func() {
		mw.logger.Log("method", "GetById", "id", id, "pet", pet, "err", err)
	}()

	return mw.next.GetById(ctx, id)
}
