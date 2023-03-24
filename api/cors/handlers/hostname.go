package corsHandler

import (
	"errors"

	"github.com/gin-gonic/gin"
	domainEntity "github.com/ruriazz/gopen-api/api/cors/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"
	"github.com/ruriazz/gopen-api/src/constants"
	responseHelper "github.com/ruriazz/gopen-api/src/helpers/response"
	"github.com/ruriazz/gopen-api/src/models"
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
		Data:    h.Serializers.Hostname().DefaultConsumerInfoV1(*result),
	})
}

func (h HostnameHandler) GetInfoV1(context *gin.Context) {
	data, _ := context.Get(constants.VAR_CONSUMER_DATA)
	if data == nil {
		return
	}

	responseHelper.JSON(responseHelper.FieldsV1{
		Context: context,
		Data:    h.Serializers.Hostname().DefaultConsumerInfoV1(data.(models.Consumer)),
	})
}

func (h HostnameHandler) NewChallenge(context *gin.Context) {

}
