package masterDataSerializer

import (
	"fmt"

	domainEntities "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	"github.com/ruriazz/gopen-api/src/models"
)

func (s MasterDataSerializer) IdnDistrict() domainInterface.IdnDistrictSerializers {
	return IdnDistrictSerializer{&s}
}

func (s IdnDistrictSerializer) DefaultIdnDistrictCollectionV1(dataModel []models.IdnDistrict) []domainEntities.DefaultIdnDistrictCollectionWithLongNameV1 {
	results := make([]domainEntities.DefaultIdnDistrictCollectionWithLongNameV1, 0)

	for _, data := range dataModel {
		alias := ""
		if data.Alias != "" {
			alias = fmt.Sprintf(" (%s)", data.Alias)
		}
		results = append(results, domainEntities.DefaultIdnDistrictCollectionWithLongNameV1{
			Slug:     data.Slug,
			Name:     fmt.Sprintf("%s %s%s", data.DatiName, data.Name, alias),
			LongName: fmt.Sprintf("%s %s, %s", data.DatiName, data.Name, data.IdnProvince.Name),
		})
	}

	return results
}

func (s IdnDistrictSerializer) DefaultIdnDistrictDetailV1(dataModel models.IdnDistrict) domainEntities.DefaultIdnDistrictDetailV1 {
	return domainEntities.DefaultIdnDistrictDetailV1{
		Slug:        dataModel.Slug,
		Code:        dataModel.Code,
		DatiName:    dataModel.DatiName,
		Name:        dataModel.Name,
		Alias:       dataModel.Alias,
		IdnProvince: s.IdnProvince().DefaultDetailV1(dataModel.IdnProvince),
	}
}

func (s IdnDistrictSerializer) DefaultIdnSubdistrictCollectionV1(dataModel []models.IdnSubdistrict) []domainEntities.DefaultIdnSubdistrictCollectionV1 {
	results := make([]domainEntities.DefaultIdnSubdistrictCollectionV1, 0)

	for _, data := range dataModel {
		alias := ""
		if data.Alias != "" {
			alias = fmt.Sprintf(" (%s)", data.Alias)
		}

		results = append(results, domainEntities.DefaultIdnSubdistrictCollectionV1{
			Slug: data.Slug,
			Name: fmt.Sprintf("%s%s", data.Name, alias),
		})
	}

	return results
}
