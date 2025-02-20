package transport

import (
	"github.com/dreadster3/pawcare/services/account/internal/pet/endpoint"
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
		kithttp.ServerErrorEncoder(common.ErrorEncoder(err2Status)),
	}
	options = append(options, common.HTTPLoggingServerOptions(logger)...)

	authenticatedOpts := append(options, kithttp.ServerBefore(kitjwt.HTTPToContext()))

	createHandler := kithttp.NewServer(
		endpoints.CreateEndpoint,
		common.DecodeJSONRequest[endpoint.CreateRequest],
		common.EncodeResponse(err2Status),
		authenticatedOpts...,
	)

	getAllHandler := kithttp.NewServer(
		endpoints.GetAllEndpoint,
		common.DecodeNoBodyRequest,
		common.EncodeResponse(err2Status),
		authenticatedOpts...,
	)

	getByIdHandler := kithttp.NewServer(
		endpoints.GetByIdEndpoint,
		common.DecodePathParameters[endpoint.GetByIdRequest],
		common.EncodeResponse(err2Status),
		authenticatedOpts...,
	)

	router := r.PathPrefix("/pets").Subrouter()
	router.Methods("POST").Path("").Handler(createHandler)
	router.Methods("GET").Path("").Handler(getAllHandler)
	router.Methods("GET").Path("/{id}").Handler(getByIdHandler)
}

func err2Status(err error) int {
	switch err {
	default:
		return common.DefaultErr2Status(err)
	}
}
