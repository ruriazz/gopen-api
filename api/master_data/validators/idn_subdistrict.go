package masterDataValidator

import (
	"errors"

	"github.com/gin-gonic/gin"
	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
)

func (v MasterDataValidator) IdnSubdistrict() domainInterface.IdnSubdistrictValidators {
	return IdnSubdistrictValidator{&v}
}

func (v IdnSubdistrictValidator) GetCollectionParameterV1(ctx *gin.Context) (*gin.Context, error) {
	var params domainEntity.GetSubdistrictCollectionParameterV1
	if err := ctx.BindQuery(&params); err != nil {
		return ctx, err
	}

	if (params == domainEntity.GetSubdistrictCollectionParameterV1{}) || params.Limit == 0 || params.Page == 0 {
		return ctx, errors.New("")
	}

	if params.Limit > 300 {
		params.Limit = 300
	}

	ctx.Set("queries", params)
	return ctx, nil
}

func (v IdnSubdistrictValidator) GetUrbanVillageCollectionParameterV1(context *gin.Context) (*gin.Context, error) {
	var params domainEntity.GetUrbanVillageCollectionBySubdistrictParameterV1
	if err := context.BindQuery(&params); err != nil {
		return context, err
	}

	if (params == domainEntity.GetUrbanVillageCollectionBySubdistrictParameterV1{}) || params.Limit == 0 || params.Page == 0 {
		return context, errors.New("")
	}

	if params.Limit > 300 {
		params.Limit = 300
	}

	context.Set("queries", params)
	return context, nil
}
