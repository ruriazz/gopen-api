package masterDataHandler

import domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"

func (h MasterDataHandler) IdnSubdistrict() domainInterface.IdnSubdistrictHandlers {
	return IdnSubdistrictHandler{&h}
}
