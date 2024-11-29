package service

import "errors"

var (
	ErrInvalidName   = errors.New("invalid name")
	ErrInvalidUserId = errors.New("invalid user id")
	ErrInvalidDate   = errors.New("invalid date of birth")
	ErrAlreadyExists = errors.New("account already exists")
)
