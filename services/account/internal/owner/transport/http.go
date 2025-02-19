package transport

import (
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
		kithttp.ServerErrorEncoder(common.ErrorEncoder(err2status)),
		kithttp.ServerBefore(kitjwt.HTTPToContext()),
	}
	options = append(options, common.HTTPLoggingServerOptions(logger)...)

	createHandler := kithttp.NewServer(
		endpoints.CreateEndpoint,
		common.DecodeJSONRequest[endpoint.CreateRequest],
		common.EncodeResponse(err2status),
		options...,
	)

	getHandler := kithttp.NewServer(
		endpoints.GetEndpoint,
		common.DecodeNoBodyRequest,
		common.EncodeResponse(err2status),
		options...,
	)

	router := r.PathPrefix("/owners").Subrouter()
	router.Methods("POST").Path("").Handler(createHandler)
	router.Methods("GET").Path("").Handler(getHandler)
}

func err2status(err error) int {
	switch err {
	case service.ErrInvalidDate:
		return http.StatusBadRequest
	default:
		return common.DefaultErr2Status(err)
	}
}
