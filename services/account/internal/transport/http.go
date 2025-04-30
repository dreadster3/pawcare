package transport

import (
	"net/http"

	ownerendpoint "github.com/dreadster3/pawcare/services/account/internal/owner/endpoint"
	ownertransport "github.com/dreadster3/pawcare/services/account/internal/owner/transport"
	petendpoint "github.com/dreadster3/pawcare/services/account/internal/pet/endpoint"
	pettransport "github.com/dreadster3/pawcare/services/account/internal/pet/transport"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func MakeHTTPServer(ownerEndpoints ownerendpoint.Set, petEndpoints petendpoint.Set, logger *zap.Logger) http.Handler {
	router := mux.NewRouter()
	apiGroup := router.PathPrefix("/api/v1").Subrouter()

	ownertransport.RegisterHTTPRoutes(apiGroup, ownerEndpoints, logger)
	pettransport.RegisterHTTPRoutes(apiGroup, petEndpoints, logger)

	return router
}
