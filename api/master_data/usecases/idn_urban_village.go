package masterDataUsecase

import (
	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	paginationHelper "github.com/ruriazz/gopen-api/helpers/pagination"
	"github.com/ruriazz/gopen-api/models"
)

func (uc MasterDataUsecase) IdnUrbanVillage() domainInterface.IdnUrbanVillageUsecases {
	return IdnUrbanVillageUsecase{&uc}
}

func (uc IdnUrbanVillageUsecase) GetCollectionV1(queries domainEntity.GetUrbanVillageCollectionParameterV1) ([]models.IdnUrbanVillage, *paginationHelper.PaginationV1, error) {
	results, pagination, err := uc.Repositories.IdnUrbanVillage().CollectionV1(queries, true)
	if err != nil {
		return nil, nil, err
	}

	return results, pagination, nil
}

func (uc IdnUrbanVillageUsecase) GetDetailV1(slug string) (*models.IdnUrbanVillage, error) {
	result, err := uc.Repositories.IdnUrbanVillage().DetailV1(models.IdnUrbanVillage{Slug: slug})
	if err != nil {
		return nil, err
	}

	return result, nil
}
