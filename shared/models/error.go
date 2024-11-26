package models

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Error: err.Error(),
	}
}

func NewErrorResponseString(err string) ErrorResponse {
	return ErrorResponse{
		Error: err,
	}
}
