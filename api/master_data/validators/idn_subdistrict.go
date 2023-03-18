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

	ctx.Set("queries", params)
	return ctx, nil
}
