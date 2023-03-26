package corsSerializer

import (
	"fmt"

	domainEntity "github.com/ruriazz/gopen-api/api/cors/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"
	"github.com/ruriazz/gopen-api/src/constants"
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

func (s HostnameSerializer) DefaultCreateDNSChallengeV1(dataModel models.ConsumerChallenge) domainEntity.DefaultCreateDNSChallengeV1 {
	challengeType := constants.DNS
	if dataModel.ChallengeType == constants.ACME.String() {
		challengeType = constants.ACME
	}

	return domainEntity.DefaultCreateDNSChallengeV1{
		ChallengeType: challengeType.LongString(),
		RecordData: domainEntity.RecordData{
			Target: dataModel.Consumer.Hostname,
			Type:   "TXT",
			Data:   dataModel.ChallengeValue,
		},
		ExpiredAt: dataModel.ExpiredAt,
	}
}

func (s HostnameSerializer) DefaultCreateACMEChallengeV1(dataModel models.ConsumerChallenge) domainEntity.DefaultCreateACMEChallengeV1 {
	challengeType := constants.DNS
	if dataModel.ChallengeType == constants.ACME.String() {
		challengeType = constants.ACME
	}

	return domainEntity.DefaultCreateACMEChallengeV1{
		ChallengeType: challengeType.LongString(),
		ACMEFile:      fmt.Sprintf(constants.ACME_LOCATION, dataModel.Consumer.Hostname, dataModel.ChallengeValue),
		ExpiredAt:     dataModel.ExpiredAt,
	}
}
