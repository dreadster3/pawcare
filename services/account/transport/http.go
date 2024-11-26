package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dreadster3/pawcare/services/account/endpoint"
	"github.com/dreadster3/pawcare/shared/models"
	"github.com/gin-gonic/gin"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
	"github.com/golang-jwt/jwt/v4"
)

func MakeHTTPHandler(endpoints endpoint.Set, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	authenticatedOpts := append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))

	createOwnerHandler := kithttp.NewServer(
		endpoints.CreateAccountEndpoint,
		decodeCreateOwnerRequest,
		encodeResponse,
		authenticatedOpts...,
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
	case kitjwt.ErrTokenExpired, kitjwt.ErrTokenContextMissing, kitjwt.ErrTokenInvalid, kitjwt.ErrTokenMalformed, kitjwt.ErrTokenNotActive, jwt.ErrSignatureInvalid:
		w.WriteHeader(http.StatusUnauthorized)

	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(models.NewErrorResponse(err))
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
