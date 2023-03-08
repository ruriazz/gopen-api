package masterDataRepository

import (
	domainInterface "github.com/ruriazz/gopen-api/openapi/master-data/domain/interface"
	"github.com/ruriazz/gopen-api/package/manager"
)

type IdnProvinceRepository struct {
	manager.Manager
}

func NewIdnProvinceRepository(mgr manager.Manager) domainInterface.IdnProvinceRepositories {
	return IdnProvinceRepository{
		Manager: mgr,
	}
}
