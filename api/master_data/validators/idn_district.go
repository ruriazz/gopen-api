package masterDataValidator

import (
	"errors"

	"github.com/gin-gonic/gin"
	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
)

func (v MasterDataValidator) IdnDistrict() domainInterface.IdnDistrictValidators {
	return IdnDistrictValidator{&v}
}

func (v IdnDistrictValidator) GetCollectionParameterV1(ctx *gin.Context) (*gin.Context, error) {
	var params domainEntity.GetDistrictCollectionParameterV1
	if err := ctx.BindQuery(&params); err != nil {
		return ctx, err
	}

	if (params == domainEntity.GetDistrictCollectionParameterV1{}) || params.Limit == 0 || params.Page == 0 {
		return ctx, errors.New("")
	}

	ctx.Set("queries", params)
	return ctx, nil
}

func (v IdnDistrictValidator) GetSubdistrictCollectionParameterV1(ctx *gin.Context) (*gin.Context, error) {
	var params domainEntity.GetSubdistrictCollectionByDistrictParameterV1
	if err := ctx.BindQuery(&params); err != nil {
		return ctx, err
	}

	if (params == domainEntity.GetSubdistrictCollectionByDistrictParameterV1{}) || params.Limit == 0 || params.Page == 0 {
		return ctx, errors.New("")
	}

	ctx.Set("queries", params)
	return ctx, nil
}
