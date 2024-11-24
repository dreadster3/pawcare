package service

import (
	"time"

	"github.com/go-kit/log"
)

type ProfileService interface {
	CreateAccount(name string, dateOfBirth time.Time) (string, error)
}

type profileService struct{}

func NewProfileService(logger log.Logger) ProfileService {
	var svc ProfileService
	svc = &profileService{}
	svc = LoggingMiddleware(logger)(svc)
	return svc
}

func (svc *profileService) CreateAccount(name string, dateOfBirth time.Time) (string, error) {
	return name, nil
}
