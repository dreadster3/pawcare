package service

import (
	"context"
	"time"

	"github.com/dreadster3/pawcare/services/auth"
	"github.com/go-kit/log"
)

type IProfileService interface {
	CreateAccount(ctx context.Context, userId string, name string, dateOfBirth time.Time) (string, error)
}

type profileService struct {
	userService auth.IUserService
}

func NewProfileService(userService auth.IUserService, logger log.Logger) IProfileService {
	var svc IProfileService
	svc = &profileService{
		userService: userService,
	}
	svc = newLoggingMiddleware(logger)(svc)
	return svc
}

func (svc *profileService) CreateAccount(ctx context.Context, userId string, name string, dateOfBirth time.Time) (string, error) {
	user, err := svc.userService.GetById(userId)
	if err != nil {
		return "", err
	}

	return string(user.Id), nil
}
