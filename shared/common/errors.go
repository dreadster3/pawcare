package common

import "errors"

var (
	ErrCastRequest   = errors.New("cannot cast request")
	ErrCastResponse  = errors.New("cannot cast response")
	ErrParsingClaims = errors.New("error parsing claims")
	ErrUnauthorized  = errors.New("unauthorized")
)
