package masterDataHandler

import (
	"github.com/gin-gonic/gin"
	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	responseHelper "github.com/ruriazz/gopen-api/src/helpers/response"
)

func (h MasterDataHandler) IdnUrbanVillage() domainInterface.IdnUrbanVillageHandlers {
	return IdnUrbanVillageHandler{&h}
}

func (h IdnUrbanVillageHandler) GetCollectionV1(context *gin.Context) {
	context, err := h.Validators.IdnUrbanVillage().GetUrbanVillageCollectionParameterV1(context)
	if err != nil {
		responseHelper.JSON(responseHelper.FieldsV1{
			Context:  context,
			MetaCode: "E0001",
		})
		return
	}

	queries, _ := context.Get("queries")
	results, pagination, err := h.Usecases.IdnUrbanVillage().GetCollectionV1(queries.(domainEntity.GetUrbanVillageCollectionParameterV1))
	if err != nil {
		responseHelper.JSON(responseHelper.FieldsV1{
			Context:  context,
			MetaCode: "E0003",
		})
		return
	}

	responseHelper.JSON(responseHelper.FieldsV1{
		Context:    context,
		Data:       h.Serializers.IdnUrbanVillage().DefaultIdnUrbanVillageCollectionWithLongNameV1(results),
		Pagination: pagination,
	})
}

func (h IdnUrbanVillageHandler) GetDetailV1(context *gin.Context) {
	slug := context.Param("slug")
	result, err := h.Usecases.IdnUrbanVillage().GetDetailV1(slug)
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
		Data:    h.Serializers.IdnUrbanVillage().DefaultIdnUrbanVillageDetailV1(*result),
	})
}
