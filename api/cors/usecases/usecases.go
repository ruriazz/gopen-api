package corsUsecase

import (
	domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"
	"github.com/ruriazz/gopen-api/package/manager"
)

type CorsUsecase struct {
	Manager      manager.Manager
	Repositories domainInterface.CorsRepositories
}

type HostnameUsecase struct{ *CorsUsecase }

func NewCorsUsecase(manager manager.Manager, repositories domainInterface.CorsRepositories) domainInterface.CorsUsecases {
	return CorsUsecase{
		Manager:      manager,
		Repositories: repositories,
	}
}
