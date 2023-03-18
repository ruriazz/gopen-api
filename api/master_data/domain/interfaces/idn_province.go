package masterDataDomainInterface

import (
	"github.com/gin-gonic/gin"
	domainEntities "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	paginationHelper "github.com/ruriazz/gopen-api/helpers/pagination"
	"github.com/ruriazz/gopen-api/models"
)

type IdnProvinceHandlers interface {
	GetCollectionV1(context *gin.Context)
	GetDistrictCollectionV1(context *gin.Context)
	GetDetailV1(context *gin.Context)
}

type IdnProvinceUsecases interface {
	GetCollectionV1(queries domainEntities.GetProvinceCollectionParameterV1) ([]models.IdnProvince, *paginationHelper.PaginationV1, error)
	GetDistrictCollectionV1(provinceSlug string, queries domainEntities.GetDistrictCollectionByProvinceParameterV1) ([]models.IdnDistrict, *paginationHelper.PaginationV1, error)
	GetDetailV1(slug string) (*models.IdnProvince, error)
}

type IdnProvinceRepositories interface {
	CollectionV1(queries *domainEntities.GetProvinceCollectionParameterV1, pagination bool) ([]models.IdnProvince, *paginationHelper.PaginationV1, error)
	DistrictCollectionV1(model models.IdnProvince, queries *domainEntities.GetDistrictCollectionByProvinceParameterV1, withPagination bool) ([]models.IdnDistrict, *paginationHelper.PaginationV1, error)
	DetailV1(model models.IdnProvince) (*models.IdnProvince, error)
}

type IdnProvinceSerializers interface {
	DefaultIdnProvinceCollectionsV1(dataModel []models.IdnProvince) []domainEntities.DefaultIdnProvinceCollectionV1
	DefaultDistrictCollectionV1(dataModel []models.IdnDistrict) []domainEntities.DefaultIdnDistrictCollectionV1
	DefaultDetailV1(dataModel models.IdnProvince) domainEntities.DefaultIdnProvinceDetailV1
}

type IdnProvinceValidators interface {
	GetCollectionParameterV1(ctx *gin.Context) (*gin.Context, error)
	GetDistrictCollectionParameterV1(ctx *gin.Context) (*gin.Context, error)
}
