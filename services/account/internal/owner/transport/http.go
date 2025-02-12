package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dreadster3/pawcare/services/account/internal/owner/endpoint"
	"github.com/dreadster3/pawcare/services/account/internal/owner/service"
	"github.com/dreadster3/pawcare/shared/common"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func RegisterHTTPRoutes(r *mux.Router, endpoints endpoint.Set, logger kitlog.Logger) {
	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
		kithttp.ServerBefore(kitjwt.HTTPToContext()),
	}
	options = append(options, common.HTTPLoggingServerOptions(logger)...)

	createHandler := kithttp.NewServer(
		endpoints.CreateEndpoint,
		common.DecodeJSONRequest[endpoint.CreateRequest],
		common.EncodeResponse(encodeError),
		options...,
	)

	getHandler := kithttp.NewServer(
		endpoints.GetEndpoint,
		common.DecodeNoBodyRequest,
		common.EncodeResponse(encodeError),
		options...,
	)

	router := r.PathPrefix("/owners").Subrouter()
	router.Methods("POST").Path("").Handler(createHandler)
	router.Methods("GET").Path("").Handler(getHandler)
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case service.ErrInvalidDate:
		w.WriteHeader(http.StatusBadRequest)
	default:
		common.EncodeError(ctx, err, w)
	}
	json.NewEncoder(w).Encode(common.NewErrorResponse(err))
}
