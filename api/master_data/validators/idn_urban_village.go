package masterDataValidator

import (
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
)

func (v MasterDataValidator) IdnUrbanVillage() domainInterface.IdnUrbanVillageValidators {
	return IdnUrbanVillageValidator{&v}
}
