package repository

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrInvalidId      = errors.New("invalid id")
)
