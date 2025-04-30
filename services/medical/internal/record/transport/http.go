package transport

import (
	"github.com/dreadster3/pawcare/services/medical/internal/record/endpoint"
	"github.com/dreadster3/pawcare/shared/common"
	"github.com/dreadster3/pawcare/shared/utils"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func RegisterHTTPRoutes(r *mux.Router, enpoints endpoint.Set, logger *zap.Logger) {
	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(common.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(common.ErrorEncoder(err2Status)),
		kithttp.ServerBefore(kitjwt.HTTPToContext()),
		kithttp.ServerBefore(utils.RequestIdHTTPToContext()),
	}
	options = append(options, common.HTTPLoggingServerOptions(logger)...)

	createHandler := kithttp.NewServer(
		enpoints.CreateEndpoint,
		common.DecodeJSONRequest[endpoint.CreateRequest],
		common.EncodeResponse(err2Status),
		options...,
	)

	getByPetIdHandler := kithttp.NewServer(
		enpoints.GetByPetIdEndpoint,
		common.DecodePathParameters[endpoint.GetByIdRequest],
		common.EncodeResponse(err2Status),
		options...,
	)

	getByIdHandler := kithttp.NewServer(
		enpoints.GetByIdEndpoint,
		common.DecodePathParameters[endpoint.GetByIdRequest],
		common.EncodeResponse(err2Status),
		options...,
	)

	r.Methods("GET").Path("/pets/{id}/records").Handler(getByPetIdHandler)
	r.Methods("POST").Path("/pets/{id}/records").Handler(createHandler)
	r.Methods("GET").Path("/records/{id}").Handler(getByIdHandler)
}

func err2Status(err error) int {
	switch err {
	default:
		return common.DefaultErr2Status(err)
	}
}
