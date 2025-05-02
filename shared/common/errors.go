package common

import "errors"

var (
	ErrCastRequest   = errors.New("cannot cast request")
	ErrCastResponse  = errors.New("cannot cast response")
	ErrParsingClaims = errors.New("error parsing claims")
	ErrUnauthorized  = errors.New("unauthorized")
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Error: err.Error(),
	}
}

func Err2Str(err error) string {
	if err == nil {
		return ""
	}

	return err.Error()
}
