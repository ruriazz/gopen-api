package masterDataHandler

import (
	"github.com/gin-gonic/gin"
	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	responseHelper "github.com/ruriazz/gopen-api/helpers/response"
)

func (h MasterDataHandler) IdnDistrict() domainInterface.IdnDistrictHandlers {
	return IdnDistrictHandler{&h}
}

func (h IdnDistrictHandler) GetCollectionV1(context *gin.Context) {
	context, err := h.Validators.IdnDistrict().GetCollectionParameterV1(context)
	if err != nil {
		responseHelper.JSON(responseHelper.FieldsV1{
			Context:  context,
			MetaCode: "E0001",
		})
		return
	}

	provinceSlug := context.Param("slug")
	queries, _ := context.Get("queries")
	results, pagination, err := h.Usecases.IdnDistrict().GetCollectionV1(provinceSlug, queries.(domainEntity.GetDistrictCollectionParameterV1))
	if err != nil {
		responseHelper.JSON(responseHelper.FieldsV1{
			Context:  context,
			MetaCode: "E0003",
		})
		return
	}

	if results == nil {
		responseHelper.JSON(responseHelper.FieldsV1{
			Context:  context,
			MetaCode: "S0001",
		})
		return
	}

	responseHelper.JSON(responseHelper.FieldsV1{
		Context:    context,
		Data:       h.Serializers.IdnDistrict().DefaultIdnDistrictCollectionV1(results),
		Pagination: pagination,
	})
}

func (h IdnDistrictHandler) GetDetailV1(context *gin.Context) {
	slug := context.Param("slug")
	result, err := h.Usecases.IdnDistrict().GetDetailV1(slug)
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
		Data:    h.Serializers.IdnDistrict().DefaultIdnDistrictDetailV1(*result),
	})
}
