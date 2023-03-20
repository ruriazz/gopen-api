package corsSerializer

import (
	domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"
	"github.com/ruriazz/gopen-api/package/databases"
	"github.com/ruriazz/gopen-api/package/manager"
)

type CorsSerializer struct {
	Manager   manager.Manager
	Databases databases.Database
}

type HostnameSerializer struct{ *CorsSerializer }

func NewCorsSerializer(manager manager.Manager) domainInterface.CorsSerializers {
	return CorsSerializer{
		Manager:   manager,
		Databases: *manager.Databases,
	}
}
