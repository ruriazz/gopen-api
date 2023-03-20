package corsHandler

import (
	domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"
	repositories "github.com/ruriazz/gopen-api/api/cors/repositories"
	serializers "github.com/ruriazz/gopen-api/api/cors/serializers"
	usecases "github.com/ruriazz/gopen-api/api/cors/usecases"
	validators "github.com/ruriazz/gopen-api/api/cors/validators"
	"github.com/ruriazz/gopen-api/package/manager"
)

type CorsHandler struct {
	Manager     manager.Manager
	Validators  domainInterface.CorsValidators
	Usecases    domainInterface.CorsUsecases
	Serializers domainInterface.CorsSerializers
}

type HostnameHandler struct{ *CorsHandler }

func NewCorsHandler(manager manager.Manager) domainInterface.CorsHandlers {
	repo := repositories.NewCorsRepository(manager)

	return CorsHandler{
		Manager:     manager,
		Validators:  validators.NewCorsValidator(manager, repo),
		Usecases:    usecases.NewCorsUsecase(manager, repo),
		Serializers: serializers.NewCorsSerializer(manager),
	}
}
