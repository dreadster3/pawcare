package service

import (
	"context"
	"time"

	"github.com/dreadster3/pawcare/services/account/aggregate"
	"github.com/dreadster3/pawcare/services/account/entity"
	"github.com/dreadster3/pawcare/services/account/repository"
	"github.com/dreadster3/pawcare/services/auth"
	"github.com/go-kit/log"
)

type IAccountService interface {
	CreateAccount(ctx context.Context, userId string, name string, dateOfBirth time.Time) (aggregate.Account, error)
}

type accountService struct {
	ownerRepository repository.IOwnerRepository
	userService     auth.IUserService
}

func NewAccountService(ownerRepository repository.IOwnerRepository, userService auth.IUserService, logger log.Logger) IAccountService {
	var svc IAccountService
	svc = &accountService{
		ownerRepository: ownerRepository,
		userService:     userService,
	}
	svc = newLoggingMiddleware(logger)(svc)
	return svc
}

func (svc *accountService) CreateAccount(ctx context.Context, userId string, name string, dateOfBirth time.Time) (aggregate.Account, error) {
	user, err := svc.userService.GetById(userId)
	if err != nil {
		return aggregate.Account{}, err
	}

	owner := &entity.Owner{Name: name, DateOfBirth: dateOfBirth}

	err = svc.ownerRepository.Create(ctx, user.Id, owner)
	if err != nil {
		return aggregate.Account{}, err
	}

	account := aggregate.Account{
		User:  user,
		Owner: owner,
		Pets:  []*entity.Pet{},
	}

	return account, nil
}
