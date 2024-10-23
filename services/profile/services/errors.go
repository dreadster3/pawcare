package services

import (
	"errors"

	"github.com/dreadster3/pawcare/services/profile/repository"
)

var (
	ErrProfileNotFound     = repository.ErrNotFound
	ErrInvalidId           = repository.ErrInvalidId
	ErrObjectAlreadyExists = errors.New("Object already exists")
)
