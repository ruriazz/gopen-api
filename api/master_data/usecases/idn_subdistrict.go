package masterDataUsecase

import (
	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	paginationHelper "github.com/ruriazz/gopen-api/src/helpers/pagination"
	"github.com/ruriazz/gopen-api/src/models"
)

func (uc MasterDataUsecase) IdnSubdistrict() domainInterface.IdnSubdistrictUsecases {
	return IdnSubdistrictUsecase{&uc}
}

func (uc IdnSubdistrictUsecase) GetCollectionV1(queries domainEntity.GetSubdistrictCollectionParameterV1) ([]models.IdnSubdistrict, *paginationHelper.PaginationV1, error) {
	results, pagination, err := uc.Repositories.IdnSubdistrict().CollectionV1(queries, true)
	if err != nil {
		return nil, nil, err
	}

	return results, pagination, nil
}

func (uc IdnSubdistrictUsecase) GetDetailV1(slug string) (*models.IdnSubdistrict, error) {
	result, err := uc.Repositories.IdnSubdistrict().DetailV1(models.IdnSubdistrict{Slug: slug})
	if err != nil {
		return nil, err
	}

	if result != nil {
		return result, nil
	}

	return nil, nil
}

func (uc IdnSubdistrictUsecase) GetUrbanVillageCollectionV1(slug string, queries domainEntity.GetUrbanVillageCollectionBySubdistrictParameterV1) ([]models.IdnUrbanVillage, *paginationHelper.PaginationV1, error) {
	subdistrict, err := uc.Repositories.IdnSubdistrict().DetailV1(models.IdnSubdistrict{Slug: slug})
	if err != nil {
		return nil, nil, err
	}

	if subdistrict == nil {
		return nil, nil, nil
	}

	results, pagination, err := uc.Repositories.IdnSubdistrict().UrbanVillageCollectionV1(*subdistrict, queries, true)
	if err != nil {
		return nil, nil, err
	}

	if results != nil {
		return results, pagination, nil
	}

	return nil, nil, nil
}
