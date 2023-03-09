package apiRoute

import (
	"github.com/gin-gonic/gin"
	masterDataRoute "github.com/ruriazz/gopen-api/api/master_data/domain"
	"github.com/ruriazz/gopen-api/package/manager"
)

func NewApiRoute(manager *manager.Manager) {
	manager.Server.Engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1 := manager.Server.Engine.Group("/v1")
	{
		masterDataRoute.NewRouterV1(v1.Group("/md"), manager)
	}
}
