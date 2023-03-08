package masterDataHandler

import (
	"github.com/gin-gonic/gin"
	domainInterface "github.com/ruriazz/gopen-api/openapi/master-data/domain/interface"
	"github.com/ruriazz/gopen-api/package/manager"
	responseWritter "github.com/ruriazz/gopen-api/package/response-writter"
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
	h.Manager.RW.JsonResponse(responseWritter.FieldsV1{
		Context: ctx,
	})
}
