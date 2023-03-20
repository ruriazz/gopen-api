package corsRepository

import (
	domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"
	"github.com/ruriazz/gopen-api/package/databases"
	"github.com/ruriazz/gopen-api/package/manager"
)

type CorsRepository struct {
	Manager   manager.Manager
	Databases databases.Database
}

type HostnameRepository struct{ *CorsRepository }

func NewCorsRepository(manager manager.Manager) domainInterface.CorsRepositories {
	return CorsRepository{
		Manager:   manager,
		Databases: *manager.Databases,
	}
}
