package masterDataUsecase

import (
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	"github.com/ruriazz/gopen-api/package/manager"
)

type MasterDataUsecase struct {
	Manager      manager.Manager
	Repositories domainInterface.MasterDataRepositories
}

type IdnProvinceUsecase struct{ *MasterDataUsecase }
type IdnDistrictUsecase struct{ *MasterDataUsecase }
type IdnSubdistrictUsecase struct{ *MasterDataUsecase }
type IdnUrbanVillageUsecase struct{ *MasterDataUsecase }

func NewMasterDataUsecase(manager manager.Manager, repositories domainInterface.MasterDataRepositories) domainInterface.MasterDataUsecases {
	return &MasterDataUsecase{
		Manager:      manager,
		Repositories: repositories,
	}
}
