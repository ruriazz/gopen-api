package googlePlaceSerializer

import (
	domainInterface "github.com/ruriazz/gopen-api/api/google_place/domain/interfaces"
	"github.com/ruriazz/gopen-api/package/databases"
	"github.com/ruriazz/gopen-api/package/manager"
)

type GooglePlaceSerializer struct {
	Manager   manager.Manager
	Databases databases.Database
}

func NewGooglePlaceSerializer(manager manager.Manager) domainInterface.GooglePlaceSerializers {
	return GooglePlaceSerializer{
		Manager:   manager,
		Databases: *manager.Databases,
	}
}
