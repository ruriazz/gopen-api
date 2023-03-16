package masterDataRepository

import (
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	"github.com/ruriazz/gopen-api/package/databases"
	"github.com/ruriazz/gopen-api/package/manager"
)

type MasterDataRepository struct {
	Manager   manager.Manager
	Databases databases.Database
}

type IdnProvinceRepository struct{ *MasterDataRepository }
type IdnDistrictRepository struct{ *MasterDataRepository }
type IdnSubdistrictRepository struct{ *MasterDataRepository }
type IdnUrbanVillageRepository struct{ *MasterDataRepository }

func NewMasterDataRepository(manager manager.Manager) domainInterface.MasterDataRepositories {
	return MasterDataRepository{
		Manager:   manager,
		Databases: *manager.Databases,
	}
}
