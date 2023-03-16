package masterDataHandler

import domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"

func (h MasterDataHandler) IdnUrbanVillage() domainInterface.IdnUrbanVillageHandlers {
	return IdnUrbanVillageHandler{&h}
}
