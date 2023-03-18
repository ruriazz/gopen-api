package masterDataDomainInterface

import (
	"github.com/gin-gonic/gin"
	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	paginationHelper "github.com/ruriazz/gopen-api/helpers/pagination"
	"github.com/ruriazz/gopen-api/models"
)

type IdnSubdistrictHandlers interface {
	GetCollectionV1(context *gin.Context)
	GetDetailV1(context *gin.Context)
	GetUrbanVillageCollection(context *gin.Context)
}

type IdnSubdistrictUsecases interface {
	GetCollectionV1(queries domainEntity.GetSubdistrictCollectionParameterV1) ([]models.IdnSubdistrict, *paginationHelper.PaginationV1, error)
	GetDetailV1(slug string) (*models.IdnSubdistrict, error)
}

type IdnSubdistrictRepositories interface {
	CollectionV1(queries domainEntity.GetSubdistrictCollectionParameterV1, withPagination bool) ([]models.IdnSubdistrict, *paginationHelper.PaginationV1, error)
	DetailV1(model models.IdnSubdistrict) (*models.IdnSubdistrict, error)
}

type IdnSubdistrictSerializers interface {
	DefaultCollectionV1(dataModel []models.IdnSubdistrict) []domainEntity.DefaultIdnSubdistrictCollectionWithLongNameV1
	DefaultIdnSubdistrictDetailV1(dataModel models.IdnSubdistrict) domainEntity.DefaultIdnSubdistrictDetailV1
}

type IdnSubdistrictValidators interface {
	GetCollectionParameterV1(ctx *gin.Context) (*gin.Context, error)
}
