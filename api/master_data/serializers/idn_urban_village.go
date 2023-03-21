package masterDataSerializer

import (
	"fmt"

	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	"github.com/ruriazz/gopen-api/src/models"
)

func (s MasterDataSerializer) IdnUrbanVillage() domainInterface.IdnUrbanVillageSerializers {
	return IdnUrbanVillageSerializer{&s}
}

func (s IdnUrbanVillageSerializer) DefaultIdnUrbanVillageCollectionWithLongNameV1(dataModel []models.IdnUrbanVillage) []domainEntity.DefaultIdnUrbanVillageCollectionWithLongNameV1 {
	results := make([]domainEntity.DefaultIdnUrbanVillageCollectionWithLongNameV1, 0)

	for _, data := range dataModel {
		results = append(results, domainEntity.DefaultIdnUrbanVillageCollectionWithLongNameV1{
			Slug:       data.Slug,
			Name:       data.Name,
			LongName:   fmt.Sprintf("%s, %s, %s %s, %s", data.Name, data.IdnSubdistrict.Name, data.IdnSubdistrict.IdnDistrict.DatiName, data.IdnSubdistrict.IdnDistrict.Name, data.IdnSubdistrict.IdnDistrict.IdnProvince.Name),
			PostalCode: data.PostalCode,
		})
	}

	return results
}

func (s IdnUrbanVillageSerializer) DefaultIdnUrbanVillageDetailV1(dataModel models.IdnUrbanVillage) domainEntity.DefaultIdnUrbanVillageDetailV1 {
	return domainEntity.DefaultIdnUrbanVillageDetailV1{
		Slug:           dataModel.Slug,
		Code:           dataModel.Code,
		Name:           dataModel.Name,
		Alias:          dataModel.Alias,
		PostalCode:     dataModel.PostalCode,
		IdnSubdistrict: s.IdnSubdistrict().DefaultIdnSubdistrictDetailV1(dataModel.IdnSubdistrict),
	}
}
