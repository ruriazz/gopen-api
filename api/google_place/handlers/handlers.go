package googlePlaceHandler

import (
	domainInterface "github.com/ruriazz/gopen-api/api/google_place/domain/interfaces"
	repositories "github.com/ruriazz/gopen-api/api/google_place/repositories"
	serializers "github.com/ruriazz/gopen-api/api/google_place/serializers"
	usecases "github.com/ruriazz/gopen-api/api/google_place/usecases"
	validators "github.com/ruriazz/gopen-api/api/google_place/validators"
	"github.com/ruriazz/gopen-api/package/manager"
)

type GooglePlaceHandler struct {
	Manager     manager.Manager
	Validators  domainInterface.GooglePlaceValidators
	Usecases    domainInterface.GooglePlaceUsecases
	Serializers domainInterface.GooglePlaceSerializers
}

func NewGooglePlaceHandler(manager *manager.Manager) domainInterface.GooglePlaceHandlers {
	repo := repositories.NewGooglePlaceRepository(*manager)
	return GooglePlaceHandler{
		Manager:     *manager,
		Validators:  validators.NewGooglePlaceValidator(*manager, repo),
		Usecases:    usecases.NewGooglePlaceUsecase(*manager, repo),
		Serializers: serializers.NewGooglePlaceSerializer(*manager),
	}
}
