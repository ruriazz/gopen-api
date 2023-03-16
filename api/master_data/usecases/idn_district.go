package masterDataUsecase

import (
	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	paginationHelper "github.com/ruriazz/gopen-api/helpers/pagination"
	"github.com/ruriazz/gopen-api/models"
)

func (uc MasterDataUsecase) IdnDistrict() domainInterface.IdnDistrictUsecases {
	return IdnDistrictUsecase{&uc}
}

func (uc IdnDistrictUsecase) GetCollectionV1(provinceSlug string, queries domainEntity.GetDistrictCollectionParameterV1) ([]models.IdnDistrict, *paginationHelper.PaginationV1, error) {
	province, err := uc.Repositories.IdnProvince().DetailV1(models.IdnProvince{Slug: provinceSlug})
	if err != nil {
		return nil, nil, err
	}

	if province == nil {
		return nil, nil, nil
	}

	results, pagination, err := uc.Repositories.IdnDistrict().CollectionV1(models.IdnDistrict{IdnProvince: *province}, &queries, true)
	if err != nil {
		return nil, nil, err
	}

	return results, pagination, nil
}

func (uc IdnDistrictUsecase) GetDetailV1(slug string) (*models.IdnDistrict, error) {
	result, err := uc.Repositories.IdnDistrict().DetailV1(models.IdnDistrict{Slug: slug})
	if err != nil {
		return nil, err
	}

	if result != nil {
		return result, nil
	}

	return nil, nil
}
