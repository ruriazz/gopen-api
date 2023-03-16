package masterDataRepository

import domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"

func (r MasterDataRepository) IdnUrbanVillage() domainInterface.IdnUrbanVillageRepositories {
	return IdnUrbanVillageRepository{&r}
}
