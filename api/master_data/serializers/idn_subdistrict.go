package masterDataSerializer

import (
	"fmt"

	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	"github.com/ruriazz/gopen-api/models"
)

func (s MasterDataSerializer) IdnSubdistrict() domainInterface.IdnSubdistrictSerializers {
	return IdnSubdistrictSerializer{&s}
}

func (s IdnSubdistrictSerializer) DefaultCollectionV1(dataModel []models.IdnSubdistrict) []domainEntity.DefaultIdnSubdistrictCollectionWithLongNameV1 {
	results := make([]domainEntity.DefaultIdnSubdistrictCollectionWithLongNameV1, 0)

	for _, data := range dataModel {
		alias := ""
		if data.Alias != "" {
			alias = fmt.Sprintf(" (%s)", data.Alias)
		}

		results = append(results, domainEntity.DefaultIdnSubdistrictCollectionWithLongNameV1{
			Slug:     data.Slug,
			Name:     fmt.Sprintf("%s%s", data.Name, alias),
			LongName: fmt.Sprintf("%s, %s %s, %s", data.Name, data.IdnDistrict.DatiName, data.IdnDistrict.Name, data.IdnDistrict.IdnProvince.Name),
		})
	}

	return results
}

func (s IdnSubdistrictSerializer) DefaultIdnSubdistrictDetailV1(dataModel models.IdnSubdistrict) domainEntity.DefaultIdnSubdistrictDetailV1 {
	return domainEntity.DefaultIdnSubdistrictDetailV1{
		Slug:        dataModel.Slug,
		Code:        dataModel.Code,
		Name:        dataModel.Name,
		Alias:       dataModel.Alias,
		IdnDistrict: s.IdnDistrict().DefaultIdnDistrictDetailV1(dataModel.IdnDistrict),
	}
}
