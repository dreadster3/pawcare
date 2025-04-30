package transport

import (
	"net/http"

	"github.com/dreadster3/pawcare/services/medical/internal/record/endpoint"
	"github.com/dreadster3/pawcare/services/medical/internal/record/transport"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func MakeHTTPServer(endpoints endpoint.Set, logger *zap.Logger) http.Handler {
	router := mux.NewRouter()
	apiGroup := router.PathPrefix("/api/v1").Subrouter()

	transport.RegisterHTTPRoutes(apiGroup, endpoints, logger)

	return router
}
