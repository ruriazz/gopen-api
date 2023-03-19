package masterDataDomainInterface

import (
	"github.com/gin-gonic/gin"
	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	paginationHelper "github.com/ruriazz/gopen-api/helpers/pagination"
	"github.com/ruriazz/gopen-api/models"
)

type IdnUrbanVillageHandlers interface {
	GetCollectionV1(context *gin.Context)
	GetDetailV1(context *gin.Context)
}

type IdnUrbanVillageUsecases interface {
	GetCollectionV1(queries domainEntity.GetUrbanVillageCollectionParameterV1) ([]models.IdnUrbanVillage, *paginationHelper.PaginationV1, error)
	GetDetailV1(slug string) (*models.IdnUrbanVillage, error)
}

type IdnUrbanVillageRepositories interface {
	CollectionV1(queries domainEntity.GetUrbanVillageCollectionParameterV1, withPagination bool) ([]models.IdnUrbanVillage, *paginationHelper.PaginationV1, error)
	DetailV1(model models.IdnUrbanVillage) (*models.IdnUrbanVillage, error)
}

type IdnUrbanVillageSerializers interface {
	DefaultIdnUrbanVillageCollectionWithLongNameV1(dataModel []models.IdnUrbanVillage) []domainEntity.DefaultIdnUrbanVillageCollectionWithLongNameV1
	DefaultIdnUrbanVillageDetailV1(datModel models.IdnUrbanVillage) domainEntity.DefaultIdnUrbanVillageDetailV1
}

type IdnUrbanVillageValidators interface {
	GetUrbanVillageCollectionParameterV1(context *gin.Context) (*gin.Context, error)
}
