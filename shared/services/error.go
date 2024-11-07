package services

import "errors"

var (
	ErrInvalidAuthentication = errors.New("invalid authentication")
	ErrNotFound              = errors.New("not found")
)
