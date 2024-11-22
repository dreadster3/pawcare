package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dreadster3/pawcare/services/account/endpoint"
	"github.com/gin-gonic/gin"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
)

func MakeHTTPHandler(endpoints endpoint.Set, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	createOwnerHandler := kithttp.NewServer(
		endpoints.CreateAccountEndpoint,
		decodeCreateOwnerRequest,
		encodeResponse,
		opts...,
	)

	engine := gin.Default()
	group := engine.Group("/api/v1/accounts")
	group.Handle("POST", "/", gin.WrapH(createOwnerHandler))

	return engine.Handler()
}

func decodeCreateOwnerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
