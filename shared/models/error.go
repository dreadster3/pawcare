package models

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error     string `json:"error"`
	RequestId string `json:"request_id"`
}

type errorResponse struct {
	error     error
	requestId string
}

func NewErrorResponse(ctx *gin.Context, err error) errorResponse {
	requestId := ctx.GetString("request_id")

	return errorResponse{error: err, requestId: requestId}
}

func NewErrorResponseString(ctx *gin.Context, err string) errorResponse {
	return NewErrorResponse(ctx, errors.New(err))
}

func NewInternalErrorResponse(ctx *gin.Context) errorResponse {
	return NewErrorResponseString(ctx, "Internal Server Error")
}

func (e errorResponse) MarshalJSON() ([]byte, error) {
	rsp := ErrorResponse{
		Error:     e.error.Error(),
		RequestId: e.requestId,
	}

	return json.Marshal(rsp)
}
