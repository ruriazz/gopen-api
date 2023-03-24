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

type HostnameRepositories interface {
	SingleConsumerV1(model models.Consumer) (*models.Consumer, error)
	CreateOneV1(model models.Consumer) (*models.Consumer, error)
}

type HostnameSerializers interface {
	DefaultConsumerInfoV1(model models.Consumer) domainEntity.DefaultConsumerInfoV1
}

type HostnameValidators interface {
	RegisterV1(context *gin.Context) (*gin.Context, error)
}
