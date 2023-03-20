package googlePlaceUsecase

import (
	domainInterface "github.com/ruriazz/gopen-api/api/google_place/domain/interfaces"
	"github.com/ruriazz/gopen-api/package/manager"
)

type GooglePlaceUsecase struct {
	Manager      manager.Manager
	Repositories domainInterface.GooglePlaceRepositories
}

func NewGooglePlaceUsecase(manager manager.Manager, repositories domainInterface.GooglePlaceRepositories) domainInterface.GooglePlaceUsecases {
	return GooglePlaceUsecase{
		Manager:      manager,
		Repositories: repositories,
	}
}
