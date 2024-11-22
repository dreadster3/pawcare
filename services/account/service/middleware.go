package service

import (
	"time"

	kitlog "github.com/go-kit/log"
)

type middleware func(ProfileService) ProfileService

func LoggingMiddleware(logger kitlog.Logger) middleware {
	return func(next ProfileService) ProfileService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger kitlog.Logger
	next   ProfileService
}

func (mw loggingMiddleware) CreateAccount(name string, dateOfBirth time.Time) (id string, err error) {
	defer func() {
		mw.logger.Log("method", "CreateOwner", "name", name, "dob", dateOfBirth, "id", id, "err", err)
	}()

	return mw.next.CreateAccount(name, dateOfBirth)
}
