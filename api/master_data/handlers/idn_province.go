package masterDataHandler

import (
	"github.com/gin-gonic/gin"
	domainInterface "github.com/ruriazz/gopen-api/api/master_data/domain/interfaces"
	responseHelper "github.com/ruriazz/gopen-api/helpers/response"
	"github.com/ruriazz/gopen-api/package/manager"
)

type IdnProvinceHandler struct {
	Manager  manager.Manager
	Usecases domainInterface.IdnProvinceUsecases
}

func NewIdnProvinceHandler(mgr manager.Manager) domainInterface.IdnProvinceHandlers {
	return &IdnProvinceHandler{
		Manager: mgr,
	}
}

func (h IdnProvinceHandler) GetCollections(ctx *gin.Context) {
	responseHelper.JSON(responseHelper.FieldsV1{
		Context: ctx,
	})
}
