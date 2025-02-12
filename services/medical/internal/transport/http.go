package transport

import (
	"net/http"

	"github.com/dreadster3/pawcare/services/medical/internal/record/endpoint"
	"github.com/dreadster3/pawcare/services/medical/internal/record/transport"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func MakeHTTPServer(endpoints endpoint.Set, logger log.Logger) http.Handler {
	router := mux.NewRouter()
	apiGroup := router.PathPrefix("/api/v1").Subrouter()

	transport.RegisterHTTPRoutes(apiGroup, endpoints, logger)

	return router
}
