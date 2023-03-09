package masterDataRoute

import (
	"github.com/gin-gonic/gin"
	handler "github.com/ruriazz/gopen-api/api/master_data/handlers"
	"github.com/ruriazz/gopen-api/package/manager"
)

func NewRouterV1(router *gin.RouterGroup, manager *manager.Manager) {
	idnProvince := router.Group("/idn-province")
	{
		idnProvinceHandler := handler.NewIdnProvinceHandler(*manager)

		idnProvince.Handle("GET", "", idnProvinceHandler.GetCollections)
	}
}
