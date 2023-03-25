package corsRepository

import (
	domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"
	"github.com/ruriazz/gopen-api/package/databases"
	"github.com/ruriazz/gopen-api/package/logger"
	"github.com/ruriazz/gopen-api/package/manager"
	"github.com/ruriazz/gopen-api/src/constants"
)

type CorsRepository struct {
	Manager   manager.Manager
	Logger    logger.ExecutionLog
	Databases databases.Database
}

type HostnameRepository struct{ *CorsRepository }

func NewCorsRepository(manager manager.Manager) domainInterface.CorsRepositories {
	return CorsRepository{
		Manager:   manager,
		Logger:    logger.NewExecutionLog(constants.REPOSITORY_MODULE, "corsRepository", "CorsRepository"),
		Databases: *manager.Databases,
	}
}
