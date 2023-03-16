package masterDataUsecase

import domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"

func (uc MasterDataUsecase) IdnUrbanVillage() domainInterface.IdnUrbanVillageUsecases {
	return IdnUrbanVillageUsecase{&uc}
}
