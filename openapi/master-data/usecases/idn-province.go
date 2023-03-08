package masterDataUsecase

import (
	domainInterface "github.com/ruriazz/gopen-api/openapi/master-data/domain/interface"
	"github.com/ruriazz/gopen-api/package/manager"
)

type IdnProvinceUsecase struct {
	Manager      manager.Manager
	Repositories domainInterface.IdnProvinceRepositories
}

func NewIdnProvinceUsecase(mgr manager.Manager) domainInterface.IdnProvinceUsecases {
	return IdnProvinceUsecase{
		Manager: mgr,
	}
}

func (_self IdnProvinceUsecase) Test() {
}
