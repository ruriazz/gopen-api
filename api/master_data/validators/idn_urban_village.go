package masterDataValidator

import (
	"errors"

	"github.com/gin-gonic/gin"
	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
)

func (v MasterDataValidator) IdnUrbanVillage() domainInterface.IdnUrbanVillageValidators {
	return IdnUrbanVillageValidator{&v}
}

func (v IdnUrbanVillageValidator) GetUrbanVillageCollectionParameterV1(context *gin.Context) (*gin.Context, error) {
	var params domainEntity.GetUrbanVillageCollectionParameterV1
	if err := context.BindQuery(&params); err != nil {
		return context, err
	}

	if (params == domainEntity.GetUrbanVillageCollectionParameterV1{}) || params.Limit == 0 || params.Page == 0 {
		return context, errors.New("")
	}

	if params.Limit > 300 {
		params.Limit = 300
	}

	context.Set("queries", params)
	return context, nil
}
