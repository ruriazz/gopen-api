package corsValidator

import (
	domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"
	"github.com/ruriazz/gopen-api/package/manager"
)

type CorsValidator struct {
	Manager      manager.Manager
	Repositories domainInterface.CorsRepositories
}

func NewCorsValidator(manager manager.Manager, repositories domainInterface.CorsRepositories) domainInterface.CorsValidators {
	return CorsValidator{
		Manager:      manager,
		Repositories: repositories,
	}
}
