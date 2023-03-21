package corsDomainInterface

import (
	"github.com/gin-gonic/gin"
	domainEntity "github.com/ruriazz/gopen-api/api/cors/domain/entities"
	"github.com/ruriazz/gopen-api/src/models"
)

type HostnameHandlers interface {
	GetInfoV1(context *gin.Context)
	RegisterV1(context *gin.Context)
	NewChallenge(context *gin.Context)
}

type HostnameUsecases interface {
	RegisterV1(domainEntity.RegisterDataV1) (*models.Consumer, error)
}

type HostnameRepositories interface{}

type HostnameSerializers interface{}

type HostnameValidators interface {
	RegisterV1(context *gin.Context) (*gin.Context, error)
}
