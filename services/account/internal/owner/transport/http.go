package transport

import (
	"net/http"

	"github.com/dreadster3/pawcare/services/account/internal/owner/endpoint"
	"github.com/dreadster3/pawcare/services/account/internal/owner/service"
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
		kithttp.ServerErrorEncoder(common.ErrorEncoder(err2status)),
		kithttp.ServerBefore(kitjwt.HTTPToContext()),
		kithttp.ServerBefore(utils.RequestIdHTTPToContext()),
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
