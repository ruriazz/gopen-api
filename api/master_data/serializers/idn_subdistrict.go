package masterDataSerializer

import (
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
)

func (s MasterDataSerializer) IdnSubdistrict() domainInterface.IdnSubdistrictSerializers {
	return IdnSubdistrictSerializer{&s}
}
