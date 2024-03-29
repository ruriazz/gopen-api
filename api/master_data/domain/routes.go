package masterDataRoute

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/ruriazz/gopen-api/api/master_data/handlers"
	"github.com/ruriazz/gopen-api/package/manager"
)

func NewRouterV1(router *gin.RouterGroup, manager *manager.Manager) {
	masterDataHandler := handlers.NewMasterDataHandler(*manager)

	idnProvince := router.Group("/idn-province")
	{
		idnProvince.Handle("GET", "", masterDataHandler.IdnProvince().GetCollectionV1)
		idnProvince.Handle("GET", ":slug", masterDataHandler.IdnProvince().GetDetailV1)
		idnProvince.Handle("GET", ":slug/districts", masterDataHandler.IdnProvince().GetDistrictCollectionV1)
	}

	idnDistrict := router.Group("/idn-district")
	{
		idnDistrict.Handle("GET", "", masterDataHandler.IdnDistrict().GetCollectionV1)
		idnDistrict.Handle("GET", ":slug", masterDataHandler.IdnDistrict().GetDetailV1)
		idnDistrict.Handle("GET", ":slug/subdistricts", masterDataHandler.IdnDistrict().GetSubdistrictCollectionV1)
	}

	idnSubdistrict := router.Group("/idn-subdistrict")
	{
		idnSubdistrict.Handle("GET", "", masterDataHandler.IdnSubdistrict().GetCollectionV1)
		idnSubdistrict.Handle("GET", ":slug", masterDataHandler.IdnSubdistrict().GetDetailV1)
		idnSubdistrict.Handle("GET", ":slug/urban-villages", masterDataHandler.IdnSubdistrict().GetUrbanVillageCollectionV1)
	}

	idnUrbanVillage := router.Group("/idn-urban-village")
	{
		idnUrbanVillage.Handle("GET", "", masterDataHandler.IdnUrbanVillage().GetCollectionV1)
		idnUrbanVillage.Handle("GET", ":slug", masterDataHandler.IdnUrbanVillage().GetDetailV1)
	}
}
