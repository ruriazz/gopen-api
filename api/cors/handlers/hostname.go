package corsHandler

import (
	"errors"

	"github.com/gin-gonic/gin"
	domainEntity "github.com/ruriazz/gopen-api/api/cors/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"
	responseHelper "github.com/ruriazz/gopen-api/src/helpers/response"
)

func (h CorsHandler) Hostname() domainInterface.HostnameHandlers {
	return HostnameHandler{&h}
}

func (h HostnameHandler) RegisterV1(context *gin.Context) {
	context, err := h.Validators.Hostname().RegisterV1(context)
	if err != nil {
		responseHelper.JSON(responseHelper.FieldsV1{
			Context: context,
			Error:   errors.New("E0002"),
		})
		return
	}

	registerData, _ := context.Get("data")
	result, err := h.Usecases.Hostname().RegisterV1(registerData.(domainEntity.RegisterDataV1))
	if err != nil {
		responseHelper.JSON(responseHelper.FieldsV1{
			Context: context,
			Error:   err,
		})
		return
	}

	responseHelper.JSON(responseHelper.FieldsV1{
		Context: context,
		Data:    result,
	})
}

func (h HostnameHandler) GetInfoV1(context *gin.Context) {

}

func (h HostnameHandler) NewChallenge(context *gin.Context) {

}
