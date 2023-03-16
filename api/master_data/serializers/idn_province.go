package masterDataSerializer

import (
	domainEntities "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	"github.com/ruriazz/gopen-api/models"
)

func (s MasterDataSerializer) IdnProvince() domainInterface.IdnProvinceSerializers {
	return IdnProvinceSerializer{&s}
}

func (s IdnProvinceSerializer) DefaultIdnProvinceCollectionsV1(dataModel []models.IdnProvince) []domainEntities.DefaultIdnProvinceCollectionV1 {
	results := make([]domainEntities.DefaultIdnProvinceCollectionV1, 0)

	for _, data := range dataModel {
		results = append(results, domainEntities.DefaultIdnProvinceCollectionV1{
			Slug:  data.Slug,
			Name:  data.Name,
			Image: data.Image,
		})
	}

	return results
}

func (s IdnProvinceSerializer) DefaultDetailV1(dataModel models.IdnProvince) domainEntities.DefaultIdnProvinceDetailV1 {
	return domainEntities.DefaultIdnProvinceDetailV1{
		Slug:           dataModel.Slug,
		Code:           dataModel.Code,
		IsoCode:        dataModel.IsoCode,
		Name:           dataModel.Name,
		GeographicArea: dataModel.GeographicArea,
		Capital:        dataModel.Capital,
		Image:          dataModel.Image,
	}
}
