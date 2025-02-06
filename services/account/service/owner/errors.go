package ownerservice

import "errors"

var (
	ErrInvalidDate   = errors.New("invalid date of birth")
	ErrAlreadyExists = errors.New("account already exists")
)
