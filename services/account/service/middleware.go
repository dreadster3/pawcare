package service

import (
	"context"
	"strings"
	"time"

	"github.com/dreadster3/pawcare/services/account/aggregate"
	kitlog "github.com/go-kit/log"
)

type middleware func(IAccountService) IAccountService

func newLoggingMiddleware(logger kitlog.Logger) middleware {
	return func(next IAccountService) IAccountService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger kitlog.Logger
	next   IAccountService
}

func (mw loggingMiddleware) CreateAccount(ctx context.Context, userId string, name string, dateOfBirth time.Time) (account aggregate.Account, err error) {
	defer func() {
		mw.logger.Log("method", "CreateAccount", "userId", userId, "name", name, "dob", dateOfBirth, "account", account, "err", err)
	}()

	return mw.next.CreateAccount(ctx, userId, name, dateOfBirth)
}

func newValidationMiddleware() middleware {
	return func(next IAccountService) IAccountService {
		return validationMiddleware{next}
	}
}

type validationMiddleware struct {
	next IAccountService
}

func (mw validationMiddleware) CreateAccount(ctx context.Context, userId string, name string, dateOfBirth time.Time) (account aggregate.Account, err error) {
	userId = strings.TrimSpace(userId)
	name = strings.TrimSpace(name)

	if name == "" {
		return aggregate.Account{}, ErrInvalidName
	}

	if userId == "" {
		return aggregate.Account{}, ErrInvalidUserId
	}

	if time.Now().After(dateOfBirth) || time.Now().AddDate(-150, 0, 0).Before(dateOfBirth) {
		return aggregate.Account{}, ErrInvalidDate
	}

	return mw.next.CreateAccount(ctx, userId, name, dateOfBirth)
}
