package repository

import "errors"

var (
	ErrNotFound       = errors.New("not found error")
	ErrAlreadyCreated = errors.New("resource already exists")
)
