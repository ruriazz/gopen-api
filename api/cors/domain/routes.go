package corsRoute

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/ruriazz/gopen-api/api/cors/handlers"
	"github.com/ruriazz/gopen-api/package/manager"
	authHelper "github.com/ruriazz/gopen-api/src/helpers/auth"
)

func NewRouterV1(router *gin.RouterGroup, manager *manager.Manager) {
	corshandler := handlers.NewCorsHandler(*manager)

	hostname := router.Group("/hostname")
	{
		hostname.Handle("GET", "", authHelper.SecretKeyAuth(*manager), corshandler.Hostname().GetInfoV1)
		hostname.Handle("POST", "", corshandler.Hostname().RegisterV1)
		hostname.Handle("POST", "challenge", corshandler.Hostname().NewChallenge)
	}
}
