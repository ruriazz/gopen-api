package corsRoute

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/ruriazz/gopen-api/api/cors/handlers"
	"github.com/ruriazz/gopen-api/package/manager"
)

func NewRouterV1(router *gin.RouterGroup, manager *manager.Manager) {
	corshandler := handlers.NewCorsHandler(*manager)

	hostname := router.Group("/hostname")
	{
		hostname.Handle("GET", "", manager.Authentication.SecretKey, corshandler.Hostname().GetInfoV1)
		hostname.Handle("POST", "", corshandler.Hostname().RegisterV1)
		hostname.Handle("POST", "challenge", corshandler.Hostname().NewChallenge)
	}
}
