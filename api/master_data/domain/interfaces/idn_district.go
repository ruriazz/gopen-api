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
	GetSubdistrictCollectionV1(context *gin.Context)
}

type IdnDistrictUsecases interface {
	GetCollectionV1(queries domainEntity.GetDistrictCollectionParameterV1) ([]models.IdnDistrict, *paginationHelper.PaginationV1, error)
	GetDetailV1(slug string) (*models.IdnDistrict, error)
	GetSubdistrictCollectionV1(slug string, queries domainEntity.GetSubdistrictCollectionByDistrictParameterV1) ([]models.IdnSubdistrict, *paginationHelper.PaginationV1, error)
}

type IdnDistrictRepositories interface {
	CollectionV1(queries *domainEntity.GetDistrictCollectionParameterV1, withPagination bool) ([]models.IdnDistrict, *paginationHelper.PaginationV1, error)
	DetailV1(model models.IdnDistrict) (*models.IdnDistrict, error)
	SubdistrictCollectionV1(model models.IdnDistrict, queries domainEntity.GetSubdistrictCollectionByDistrictParameterV1, withPagination bool) ([]models.IdnSubdistrict, *paginationHelper.PaginationV1, error)
}

type IdnDistrictSerializers interface {
	DefaultIdnDistrictCollectionV1(dataModel []models.IdnDistrict) []domainEntity.DefaultIdnDistrictCollectionWithLongNameV1
	DefaultIdnDistrictDetailV1(dataModel models.IdnDistrict) domainEntity.DefaultIdnDistrictDetailV1
	DefaultIdnSubdistrictCollectionV1(dataModel []models.IdnSubdistrict) []domainEntity.DefaultIdnSubdistrictCollectionV1
}

type IdnDistrictValidators interface {
	GetCollectionParameterV1(ctx *gin.Context) (*gin.Context, error)
	GetSubdistrictCollectionParameterV1(ctx *gin.Context) (*gin.Context, error)
}
