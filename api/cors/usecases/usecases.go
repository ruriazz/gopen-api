package corsUsecase

import (
	domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"
	"github.com/ruriazz/gopen-api/package/logger"
	"github.com/ruriazz/gopen-api/package/manager"
	"github.com/ruriazz/gopen-api/src/constants"
)

type CorsUsecase struct {
	Manager      manager.Manager
	Logger       logger.ExecutionLog
	Repositories domainInterface.CorsRepositories
}

type HostnameUsecase struct{ *CorsUsecase }

func NewCorsUsecase(manager manager.Manager, repositories domainInterface.CorsRepositories) domainInterface.CorsUsecases {
	return CorsUsecase{
		Manager:      manager,
		Logger:       logger.NewExecutionLog(constants.USECASE_MODULE, "corsUsecase", "CorsUsecase"),
		Repositories: repositories,
	}
}
