package masterDataSerializer

import (
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	"github.com/ruriazz/gopen-api/package/manager"
)

type MasterDataSerializer struct {
	Manager      manager.Manager
	Repositories domainInterface.MasterDataRepositories
}

type IdnProvinceSerializer struct{ *MasterDataSerializer }
type IdnDistrictSerializer struct{ *MasterDataSerializer }
type IdnSubdistrictSerializer struct{ *MasterDataSerializer }
type IdnUrbanVillageSerializer struct{ *MasterDataSerializer }

func NewMasterDataSerializer(manager manager.Manager, repositories domainInterface.MasterDataRepositories) domainInterface.MasterDataSerializers {
	return &MasterDataSerializer{
		Manager:      manager,
		Repositories: repositories,
	}
}
