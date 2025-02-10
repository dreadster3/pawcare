package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dreadster3/pawcare/services/account/pet/endpoint"
	"github.com/dreadster3/pawcare/services/account/repository"
	"github.com/dreadster3/pawcare/shared/common"
	"github.com/dreadster3/pawcare/shared/models"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

func RegisterHTTPRoutes(r *mux.Router, endpoints endpoint.Set, logger kitlog.Logger) {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	authenticatedOpts := append(opts, kithttp.ServerBefore(kitjwt.HTTPToContext()))

	createHandler := kithttp.NewServer(
		endpoints.CreateEndpoint,
		common.DecodeJSONRequest[endpoint.CreateRequest],
		encodeResponse,
		authenticatedOpts...,
	)

	getAllHandler := kithttp.NewServer(
		endpoints.GetAllEndpoint,
		common.DecodeNoBodyRequest,
		encodeResponse,
		authenticatedOpts...,
	)

	getHandler := kithttp.NewServer(
		endpoints.GetEndpoint,
		common.DecodePathParameters[endpoint.GetRequest],
		encodeResponse,
		authenticatedOpts...,
	)

	router := r.PathPrefix("/pets").Subrouter()
	router.Methods("POST").Path("").Handler(createHandler)
	router.Methods("GET").Path("").Handler(getAllHandler)
	router.Methods("GET").Path("/{id}").Handler(getHandler)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case kitjwt.ErrTokenExpired, kitjwt.ErrTokenContextMissing, kitjwt.ErrTokenInvalid, kitjwt.ErrTokenMalformed, kitjwt.ErrTokenNotActive, jwt.ErrSignatureInvalid:
		w.WriteHeader(http.StatusUnauthorized)
	case repository.ErrAlreadyCreated:
		w.WriteHeader(http.StatusConflict)
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
