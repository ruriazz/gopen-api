package masterDataHandler

import (
	"github.com/gin-gonic/gin"
	domainEntity "github.com/ruriazz/gopen-api/api/master_data/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	responseHelper "github.com/ruriazz/gopen-api/helpers/response"
)

func (h MasterDataHandler) IdnProvince() domainInterface.IdnProvinceHandlers {
	return IdnProvinceHandler{&h}
}

func (h IdnProvinceHandler) GetCollectionV1(context *gin.Context) {
	context, err := h.Validators.IdnProvince().GetCollectionParameterV1(context)
	if err != nil {
		responseHelper.JSON(responseHelper.FieldsV1{
			Context:  context,
			MetaCode: "E0001",
		})
		return
	}

	queries, _ := context.Get("queries")
	results, pagination, err := h.Usecases.IdnProvince().GetCollectionV1(queries.(domainEntity.GetProvinceCollectionParameterV1))
	if err != nil {
		responseHelper.JSON(responseHelper.FieldsV1{
			Context:  context,
			MetaCode: "E0003",
		})
		return
	}

	result := h.Serializers.IdnProvince().DefaultIdnProvinceCollectionsV1(results)

	responseHelper.JSON(responseHelper.FieldsV1{
		Context:    context,
		Data:       result,
		Pagination: pagination,
	})
}

func (h IdnProvinceHandler) GetDistrictCollectionV1(context *gin.Context) {
	context, err := h.Validators.IdnProvince().GetDistrictCollectionParameterV1(context)
	if err != nil {
		responseHelper.JSON(responseHelper.FieldsV1{
			Context:  context,
			MetaCode: "E0001",
		})
		return
	}

	provinceSlug := context.Param("slug")
	queries, _ := context.Get("queries")
	results, pagination, err := h.Usecases.IdnProvince().GetDistrictCollectionV1(provinceSlug, queries.(domainEntity.GetDistrictCollectionByProvinceParameterV1))
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
		Data:       h.Serializers.IdnProvince().DefaultDistrictCollectionV1(results),
		Pagination: pagination,
	})
}

func (h IdnProvinceHandler) GetDetailV1(context *gin.Context) {
	slug := context.Param("slug")

	result, err := h.Usecases.IdnProvince().GetDetailV1(slug)
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
		Data:    h.Serializers.IdnProvince().DefaultDetailV1(*result),
	})
}
