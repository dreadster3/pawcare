package pet

import (
	"context"

	"go.uber.org/zap"
)

type middleware func(IPetService) IPetService

type loggingMiddleware struct {
	logger *zap.Logger
	next   IPetService
}

func newLoggingMiddleware(logger *zap.Logger) middleware {
	return func(next IPetService) IPetService {
		return &loggingMiddleware{logger, next}
	}
}

func (mw *loggingMiddleware) GetById(ctx context.Context, id PetId) (pet Pet, err error) {
	defer func() {
		mw.logger.
			Info("GetById",
				zap.String("method", "GetById"),
				zap.String("id", string(id)),
				zap.Stringer("pet", pet),
				zap.Error(err),
			)
	}()

	return mw.next.GetById(ctx, id)
}
