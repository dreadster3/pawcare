package transport

import (
	"net/http"

	ownerendpoint "github.com/dreadster3/pawcare/services/account/owner/endpoint"
	ownertransport "github.com/dreadster3/pawcare/services/account/owner/transport"
	petendpoint "github.com/dreadster3/pawcare/services/account/pet/endpoint"
	pettransport "github.com/dreadster3/pawcare/services/account/pet/transport"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
)

func MakeHTTPServer(ownerEndpoints ownerendpoint.Set, petEndpoints petendpoint.Set, logger log.Logger) http.Handler {

	engine := gin.Default()
	apiGroup := engine.Group("/api/v1")

	ownertransport.RegisterHTTPRoutes(apiGroup, ownerEndpoints, logger)
	pettransport.RegisterHTTPRoutes(apiGroup, petEndpoints, logger)

	return engine.Handler()
}
