package googlePlaceValidator

import (
	domainInterface "github.com/ruriazz/gopen-api/api/google_place/domain/interfaces"
	"github.com/ruriazz/gopen-api/package/manager"
)

type GooglePlaceValidator struct {
	Manager      manager.Manager
	Repositories domainInterface.GooglePlaceRepositories
}

func NewGooglePlaceValidator(manager manager.Manager, repositories domainInterface.GooglePlaceRepositories) domainInterface.GooglePlaceValidators {
	return GooglePlaceValidator{
		Manager:      manager,
		Repositories: repositories,
	}
}
