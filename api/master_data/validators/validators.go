package masterDataValidator

import (
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	"github.com/ruriazz/gopen-api/package/manager"
)

type MasterDataValidator struct {
	Manager      manager.Manager
	Repositories domainInterface.MasterDataRepositories
}

type IdnProvinceValidator struct{ *MasterDataValidator }
type IdnDistrictValidator struct{ *MasterDataValidator }
type IdnSubdistrictValidator struct{ *MasterDataValidator }
type IdnUrbanVillageValidator struct{ *MasterDataValidator }

func NewMasterDataValidator(manager manager.Manager, repositories domainInterface.MasterDataRepositories) domainInterface.MasterDataValidators {
	return MasterDataValidator{
		Manager:      manager,
		Repositories: repositories,
	}
}
