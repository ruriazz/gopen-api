package masterDataHandler

import (
	"github.com/gin-gonic/gin"
	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	responseHelper "github.com/ruriazz/gopen-api/helpers/response"
)

func (h MasterDataHandler) IdnSubdistrict() domainInterface.IdnSubdistrictHandlers {
	return IdnSubdistrictHandler{&h}
}

func (h IdnSubdistrictHandler) GetCollectionV1(context *gin.Context) {
	context, err := h.Validators.IdnSubdistrict().GetCollectionParameterV1(context)
	if err != nil {
		responseHelper.JSON(responseHelper.FieldsV1{
			Context:  context,
			MetaCode: "E0001",
		})
		return
	}

	queries, _ := context.Get("queries")
	results, pagination, err := h.Usecases.IdnSubdistrict().GetCollectionV1(queries.(domainEntity.GetSubdistrictCollectionParameterV1))
	if err != nil {
		responseHelper.JSON(responseHelper.FieldsV1{
			Context:  context,
			MetaCode: "E0003",
		})
		return
	}

	responseHelper.JSON(responseHelper.FieldsV1{
		Context:    context,
		Data:       h.Serializers.IdnSubdistrict().DefaultCollectionV1(results),
		Pagination: pagination,
	})
}

func (h IdnSubdistrictHandler) GetDetailV1(context *gin.Context) {
	slug := context.Param("slug")
	result, err := h.Usecases.IdnSubdistrict().GetDetailV1(slug)
	if err != nil {
		responseHelper.JSON(responseHelper.FieldsV1{
			Context:  context,
			MetaCode: "E0003",
		})
		return
	}

	if result == nil {
		responseHelper.JSON(responseHelper.FieldsV1{
			Context:  context,
			MetaCode: "S0001",
		})
		return
	}

	responseHelper.JSON(responseHelper.FieldsV1{
		Context: context,
		Data:    h.Serializers.IdnSubdistrict().DefaultIdnSubdistrictDetailV1(*result),
	})
}

func (h IdnSubdistrictHandler) GetUrbanVillageCollection(context *gin.Context) {

}
