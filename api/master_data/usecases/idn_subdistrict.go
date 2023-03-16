package masterDataUsecase

import domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"

func (uc MasterDataUsecase) IdnSubdistrict() domainInterface.IdnSubdistrictUsecases {
	return IdnSubdistrictUsecase{&uc}
}
