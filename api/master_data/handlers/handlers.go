package masterDataHandler

import (
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	repositories "github.com/ruriazz/gopen-api/api/master_data/repositories"
	serializers "github.com/ruriazz/gopen-api/api/master_data/serializers"
	usecases "github.com/ruriazz/gopen-api/api/master_data/usecases"
	validators "github.com/ruriazz/gopen-api/api/master_data/validators"
	"github.com/ruriazz/gopen-api/package/manager"
)

type MasterDataHandler struct {
	Manager     manager.Manager
	Validators  domainInterface.MasterDataValidators
	Usecases    domainInterface.MasterDataUsecases
	Serializers domainInterface.MasterDataSerializers
}

type IdnProvinceHandler struct{ *MasterDataHandler }
type IdnDistrictHandler struct{ *MasterDataHandler }
type IdnSubdistrictHandler struct{ *MasterDataHandler }
type IdnUrbanVillageHandler struct{ *MasterDataHandler }

func NewMasterDataHandler(manager manager.Manager) domainInterface.MasterDataHandlers {
	repo := repositories.NewMasterDataRepository(manager)

	return MasterDataHandler{
		Manager:     manager,
		Validators:  validators.NewMasterDataValidator(manager, repo),
		Usecases:    usecases.NewMasterDataUsecase(manager, repo),
		Serializers: serializers.NewMasterDataSerializer(manager, repo),
	}
}
