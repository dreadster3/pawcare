package common

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dreadster3/pawcare/shared/utils"
	"github.com/gorilla/mux"
)

func DecodeNoBodyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func DecodePathParameters[T any](_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	result, err := utils.MapToStruct[T](vars)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func DecodeJSONRequest[T any](_ context.Context, r *http.Request) (interface{}, error) {
	var request T
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func GRPCDecodeNoBody(_ context.Context, req interface{}) (interface{}, error) {
	return nil, nil
}
