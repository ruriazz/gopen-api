package corsDomainInterface

import (
	"github.com/gin-gonic/gin"
	domainEntity "github.com/ruriazz/gopen-api/api/cors/domain/entities"
	"github.com/ruriazz/gopen-api/src/models"
)

type HostnameHandlers interface {
	GetInfoV1(context *gin.Context)
	RegisterV1(context *gin.Context)
	CreateChallengeV1(context *gin.Context)
}

type HostnameUsecases interface {
	RegisterV1(domainEntity.RegisterDataV1) (*models.Consumer, error)
	CreateChallengeV1(consumer models.Consumer, challengeData domainEntity.CreateChallengeV1) (*models.ConsumerChallenge, error)
}

type HostnameRepositories interface {
	SingleConsumerV1(model models.Consumer) (*models.Consumer, error)
	CreateOneV1(model models.Consumer) (*models.Consumer, error)
	GetOneConsumerChallengeV1(model *models.ConsumerChallenge, preload bool) error
	CeateConsumerChallengeV1(model *models.ConsumerChallenge) error
	UpdateConsumerChallengeV1(model *models.ConsumerChallenge, newData models.ConsumerChallenge) error
}

type HostnameSerializers interface {
	DefaultConsumerInfoV1(model models.Consumer) domainEntity.DefaultConsumerInfoV1
	DefaultCreateDNSChallengeV1(dataModel models.ConsumerChallenge) domainEntity.DefaultCreateDNSChallengeV1
	DefaultCreateACMEChallengeV1(dataModel models.ConsumerChallenge) domainEntity.DefaultCreateACMEChallengeV1
}

type HostnameValidators interface {
	RegisterV1(context *gin.Context) (*gin.Context, error)
	CreateChallengeV1(context *gin.Context) error
}
