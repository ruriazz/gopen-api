package masterDataDomainInterface

import (
	"github.com/gin-gonic/gin"
	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	paginationHelper "github.com/ruriazz/gopen-api/helpers/pagination"
	"github.com/ruriazz/gopen-api/models"
)

type IdnDistrictHandlers interface {
	GetCollectionV1(context *gin.Context)
	GetDetailV1(context *gin.Context)
}

type IdnDistrictUsecases interface {
	GetCollectionV1(queries domainEntity.GetDistrictCollectionParameterV1) ([]models.IdnDistrict, *paginationHelper.PaginationV1, error)
	GetDetailV1(slug string) (*models.IdnDistrict, error)
}

type IdnDistrictRepositories interface {
	CollectionV1(queries *domainEntity.GetDistrictCollectionParameterV1, withPagination bool) ([]models.IdnDistrict, *paginationHelper.PaginationV1, error)
	DetailV1(model models.IdnDistrict) (*models.IdnDistrict, error)
}

type IdnDistrictSerializers interface {
	DefaultIdnDistrictCollectionV1(dataModel []models.IdnDistrict) []domainEntity.DefaultIdnDistrictCollectionWithLongNameV1
	DefaultIdnDistrictDetailV1(dataModel models.IdnDistrict) domainEntity.DefaultIdnDistrictDetailV1
}

type IdnDistrictValidators interface {
	GetCollectionParameterV1(ctx *gin.Context) (*gin.Context, error)
}
