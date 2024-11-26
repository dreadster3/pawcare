package auth

import "github.com/go-kit/log"

type IUserService interface {
	GetById(id string) (User, error)
}

func NewUserService(logger log.Logger) IUserService {
	var svc IUserService
	svc = &userService{}
	svc = newLoggingMiddleware(logger)(svc)
	return svc
}

type userService struct{}

func (us *userService) GetById(id string) (User, error) {
	return User{Id: UserId(id)}, nil
}
