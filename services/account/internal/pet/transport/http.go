package transport

import (
	"github.com/dreadster3/pawcare/services/account/internal/pet/endpoint"
	"github.com/dreadster3/pawcare/shared/common"
	"github.com/dreadster3/pawcare/shared/utils"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func RegisterHTTPRoutes(r *mux.Router, endpoints endpoint.Set, logger *zap.Logger) {
	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(common.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(kithttp.DefaultErrorEncoder),
		kithttp.ServerBefore(kitjwt.HTTPToContext()),
		kithttp.ServerBefore(utils.RequestIdHTTPToContext()),
	}
	options = append(options, common.HTTPLoggingServerOptions(logger)...)

	authenticatedOpts := append(options, kithttp.ServerBefore(kitjwt.HTTPToContext()))

	createHandler := kithttp.NewServer(
		endpoints.CreateEndpoint,
		common.DecodeJSONRequest[endpoint.CreateRequest],
		kithttp.EncodeJSONResponse,
		authenticatedOpts...,
	)

	getAllHandler := kithttp.NewServer(
		endpoints.GetAllEndpoint,
		common.DecodeNoBodyRequest,
		kithttp.EncodeJSONResponse,
		authenticatedOpts...,
	)

	getByIdHandler := kithttp.NewServer(
		endpoints.GetByIdEndpoint,
		common.DecodePathParameters[endpoint.GetByIdRequest],
		kithttp.EncodeJSONResponse,
		authenticatedOpts...,
	)

	router := r.PathPrefix("/pets").Subrouter()
	router.Methods("POST").Path("").Handler(createHandler)
	router.Methods("GET").Path("").Handler(getAllHandler)
	router.Methods("GET").Path("/{id}").Handler(getByIdHandler)
}
