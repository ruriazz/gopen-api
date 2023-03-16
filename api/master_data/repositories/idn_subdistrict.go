package masterDataRepository

import domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"

func (r MasterDataRepository) IdnSubdistrict() domainInterface.IdnSubdistrictRepositories {
	return IdnSubdistrictRepository{&r}
}
