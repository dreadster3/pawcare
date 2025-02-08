package common

import "errors"

var (
	ErrCastRequest   = errors.New("cannot cast request")
	ErrParsingClaims = errors.New("error parsing claims")
)
