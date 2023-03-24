package corsSerializer

import (
	"fmt"

	domainEntity "github.com/ruriazz/gopen-api/api/cors/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"
	"github.com/ruriazz/gopen-api/src/models"
)

func (s CorsSerializer) Hostname() domainInterface.HostnameSerializers {
	return HostnameSerializer{&s}
}

func (s HostnameSerializer) DefaultConsumerInfoV1(dataModel models.Consumer) domainEntity.DefaultConsumerInfoV1 {
	secretKey := fmt.Sprintf("%s:%s", dataModel.GetID(), dataModel.SecretKey)

	return domainEntity.DefaultConsumerInfoV1{
		Hostname:       dataModel.Hostname,
		MaintenerEmail: dataModel.MaintenerEmail,
		IsActive:       dataModel.IsActive,
		IsValidated:    dataModel.IsValidated,
		SecretKey:      secretKey,
	}
}
