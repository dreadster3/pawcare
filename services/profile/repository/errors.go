package repository

import "errors"

var (
	ErrInvalidId = errors.New("Invalid Id")
	ErrNotFound  = errors.New("Profile not found")
)
