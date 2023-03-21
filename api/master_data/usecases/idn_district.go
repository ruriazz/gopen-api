package masterDataUsecase

import (
	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	paginationHelper "github.com/ruriazz/gopen-api/src/helpers/pagination"
	"github.com/ruriazz/gopen-api/src/models"
)

func (uc MasterDataUsecase) IdnDistrict() domainInterface.IdnDistrictUsecases {
	return IdnDistrictUsecase{&uc}
}

func (uc IdnDistrictUsecase) GetCollectionV1(queries domainEntity.GetDistrictCollectionParameterV1) ([]models.IdnDistrict, *paginationHelper.PaginationV1, error) {
	results, pagination, err := uc.Repositories.IdnDistrict().CollectionV1(&queries, true)
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

func (uc IdnDistrictUsecase) GetSubdistrictCollectionV1(slug string, queries domainEntity.GetSubdistrictCollectionByDistrictParameterV1) ([]models.IdnSubdistrict, *paginationHelper.PaginationV1, error) {
	district, err := uc.Repositories.IdnDistrict().DetailV1(models.IdnDistrict{Slug: slug})
	if err != nil {
		return nil, nil, err
	}

	if district != nil {
		results, pagination, err := uc.Repositories.IdnDistrict().SubdistrictCollectionV1(*district, queries, true)
		if err != nil {
			return nil, nil, err
		}

		return results, pagination, nil
	}

	return nil, nil, nil
}
