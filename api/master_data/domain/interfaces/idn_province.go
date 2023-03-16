package masterDataDomainInterface

import (
	"github.com/gin-gonic/gin"
	domainEntities "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	paginationHelper "github.com/ruriazz/gopen-api/helpers/pagination"
	"github.com/ruriazz/gopen-api/models"
)

type IdnProvinceHandlers interface {
	GetCollectionV1(context *gin.Context)
	GetDetailV1(context *gin.Context)
}

type IdnProvinceUsecases interface {
	GetCollectionV1(queries domainEntities.GetProvinceCollectionParameterV1) ([]models.IdnProvince, *paginationHelper.PaginationV1, error)
	GetDetailV1(slug string) (*models.IdnProvince, error)
}

type IdnProvinceRepositories interface {
	CollectionV1(queries *domainEntities.GetProvinceCollectionParameterV1, pagination bool) ([]models.IdnProvince, *paginationHelper.PaginationV1, error)
	DetailV1(model models.IdnProvince) (*models.IdnProvince, error)
}

type IdnProvinceSerializers interface {
	DefaultIdnProvinceCollectionsV1(dataModel []models.IdnProvince) []domainEntities.DefaultIdnProvinceCollectionV1
	DefaultDetailV1(dataModel models.IdnProvince) domainEntities.DefaultIdnProvinceDetailV1
}

type IdnProvinceValidators interface {
	GetCollectionParameterV1(ctx *gin.Context) (*gin.Context, error)
}
