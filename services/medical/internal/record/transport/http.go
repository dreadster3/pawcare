package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dreadster3/pawcare/services/medical/internal/record/endpoint"
	"github.com/dreadster3/pawcare/shared/common"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func RegisterHTTPRoutes(r *mux.Router, enpoints endpoint.Set, logger log.Logger) {
	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
		kithttp.ServerBefore(kitjwt.HTTPToContext()),
	}
	options = append(options, common.HTTPLoggingServerOptions(logger)...)

	createHandler := kithttp.NewServer(
		enpoints.CreateEndpoint,
		common.DecodeJSONRequest[endpoint.CreateRequest],
		common.EncodeResponse(encodeError),
		options...,
	)

	getByPetIdHandler := kithttp.NewServer(
		enpoints.GetByPetIdEndpoint,
		common.DecodePathParameters[endpoint.GetByIdRequest],
		common.EncodeResponse(encodeError),
		options...,
	)

	getByIdHandler := kithttp.NewServer(
		enpoints.GetByIdEndpoint,
		common.DecodePathParameters[endpoint.GetByIdRequest],
		common.EncodeResponse(encodeError),
		options...,
	)

	r.Methods("GET").Path("/pets/{id}/records").Handler(getByPetIdHandler)
	router := r.PathPrefix("/records").Subrouter()
	router.Methods("POST").Path("").Handler(createHandler)
	router.Methods("GET").Path("/{id}").Handler(getByIdHandler)
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		common.EncodeError(ctx, err, w)
	}
	json.NewEncoder(w).Encode(common.NewErrorResponse(err))
}
