package masterDataSerializer

import (
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
)

func (s MasterDataSerializer) IdnUrbanVillage() domainInterface.IdnUrbanVillageSerializers {
	return IdnUrbanVillageSerializer{&s}
}
