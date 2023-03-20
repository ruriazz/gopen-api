package corsHandler

import (
	"github.com/gin-gonic/gin"
	domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"
)

func (h CorsHandler) Hostname() domainInterface.HostnameHandlers {
	return HostnameHandler{&h}
}

func (h HostnameHandler) RegisterV1(context *gin.Context) {

}

func (h HostnameHandler) GetInfoV1(context *gin.Context) {

}

func (h HostnameHandler) NewChallenge(context *gin.Context) {

}
