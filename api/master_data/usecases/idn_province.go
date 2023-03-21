package masterDataUsecase

import (
	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	paginationHelper "github.com/ruriazz/gopen-api/src/helpers/pagination"
	"github.com/ruriazz/gopen-api/src/models"
)

func (uc MasterDataUsecase) IdnProvince() domainInterface.IdnProvinceUsecases {
	return IdnProvinceUsecase{&uc}
}

func (uc IdnProvinceUsecase) GetCollectionV1(queries domainEntity.GetProvinceCollectionParameterV1) ([]models.IdnProvince, *paginationHelper.PaginationV1, error) {
	results, pagination, err := uc.Repositories.IdnProvince().CollectionV1(&queries, true)
	if err != nil {
		return nil, nil, err
	}

	return results, pagination, nil
}

func (uc IdnProvinceUsecase) GetDistrictCollectionV1(provinceSlug string, queries domainEntity.GetDistrictCollectionByProvinceParameterV1) ([]models.IdnDistrict, *paginationHelper.PaginationV1, error) {
	province, err := uc.Repositories.IdnProvince().DetailV1(models.IdnProvince{Slug: provinceSlug})
	if err != nil {
		return nil, nil, err
	}

	if province == nil {
		return nil, nil, nil
	}

	results, pagination, err := uc.Repositories.IdnProvince().DistrictCollectionV1(*province, &queries, true)
	if err != nil {
		return nil, nil, err
	}

	return results, pagination, nil
}

func (uc IdnProvinceUsecase) GetDetailV1(slug string) (*models.IdnProvince, error) {
	result, err := uc.Repositories.IdnProvince().DetailV1(models.IdnProvince{Slug: slug})
	if err != nil {
		return nil, err
	}

	return result, nil
}
