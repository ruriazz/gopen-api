package masterDataValidator

import (
	"errors"

	"github.com/gin-gonic/gin"
	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
)

func (v MasterDataValidator) IdnProvince() domainInterface.IdnProvinceValidators {
	return IdnProvinceValidator{&v}
}

func (v IdnProvinceValidator) GetCollectionParameterV1(ctx *gin.Context) (*gin.Context, error) {
	var params domainEntity.GetProvinceCollectionParameterV1

	if err := ctx.ShouldBindQuery(&params); err != nil {
		return ctx, err
	}

	if (params == domainEntity.GetProvinceCollectionParameterV1{}) || params.Limit == 0 || params.Page == 0 {
		return ctx, errors.New("")
	}

	ctx.Set("queries", params)
	return ctx, nil
}
