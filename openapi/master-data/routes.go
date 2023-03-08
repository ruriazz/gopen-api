package masterDataRoute

import (
	"github.com/gin-gonic/gin"
	masterDataHandler "github.com/ruriazz/gopen-api/openapi/master-data/handlers"
	"github.com/ruriazz/gopen-api/package/manager"
)

func NewRouterV1(rg *gin.RouterGroup, mgr *manager.Manager) {

	idnProvince := rg.Group("/idn-province")
	{
		idnProvinceHandler := masterDataHandler.NewIdnProvinceHandler(*mgr)

		idnProvince.Handle("GET", "", idnProvinceHandler.GetCollections)
	}
}

func NewRouterV2(rg *gin.RouterGroup, mgr *manager.Manager) {}
