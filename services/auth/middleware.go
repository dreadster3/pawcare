package auth

import (
	kitlog "github.com/go-kit/log"
)

type Middleware func(IUserService) IUserService

func newLoggingMiddleware(logger kitlog.Logger) Middleware {
	return func(next IUserService) IUserService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger kitlog.Logger
	next   IUserService
}

func (mw loggingMiddleware) GetById(id string) (user User, err error) {
	defer func() {
		mw.logger.Log("method", "GetById", "userId", id, "user", user, "err", err)
	}()

	return mw.next.GetById(id)
}
