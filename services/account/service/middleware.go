package service

import (
	"context"
	"strings"
	"time"

	kitlog "github.com/go-kit/log"
)

type middleware func(IProfileService) IProfileService

func newLoggingMiddleware(logger kitlog.Logger) middleware {
	return func(next IProfileService) IProfileService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger kitlog.Logger
	next   IProfileService
}

func (mw loggingMiddleware) CreateAccount(ctx context.Context, userId string, name string, dateOfBirth time.Time) (id string, err error) {
	defer func() {
		mw.logger.Log("method", "CreateAccount", "userId", userId, "name", name, "dob", dateOfBirth, "accountId", id, "err", err)
	}()

	return mw.next.CreateAccount(ctx, userId, name, dateOfBirth)
}

type validationMiddleware struct {
	next IProfileService
}

func (mw validationMiddleware) CreateAccount(ctx context.Context, userId string, name string, dateOfBirth time.Time) (id string, err error) {
	userId = strings.TrimSpace(userId)
	name = strings.TrimSpace(name)

	if name == "" {
		return "", ErrInvalidName
	}

	if userId == "" {
		return "", ErrInvalidUserId
	}

	if time.Now().After(dateOfBirth) || time.Now().AddDate(-150, 0, 0).Before(dateOfBirth) {
		return "", ErrInvalidDate
	}

	return mw.next.CreateAccount(ctx, userId, name, dateOfBirth)
}
