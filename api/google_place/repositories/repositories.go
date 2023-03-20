package googlePlaceRepository

import (
	domainInterface "github.com/ruriazz/gopen-api/api/google_place/domain/interfaces"
	"github.com/ruriazz/gopen-api/package/databases"
	"github.com/ruriazz/gopen-api/package/manager"
)

type GooglePlaceRepository struct {
	Manager   manager.Manager
	Databases databases.Database
}

func NewGooglePlaceRepository(manager manager.Manager) domainInterface.GooglePlaceRepositories {
	return GooglePlaceRepository{
		Manager:   manager,
		Databases: *manager.Databases,
	}
}
