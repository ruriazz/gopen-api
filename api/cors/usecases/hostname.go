package corsUsecase

import (
	"errors"
	"fmt"
	"time"

	domainEntity "github.com/ruriazz/gopen-api/api/cors/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"
	"github.com/ruriazz/gopen-api/package/logger"

	// captchaHelper "github.com/ruriazz/gopen-api/src/helpers/captcha"
	"github.com/ruriazz/gopen-api/src/constants"
	encryptionHelper "github.com/ruriazz/gopen-api/src/helpers/encryption"
	stringHelper "github.com/ruriazz/gopen-api/src/helpers/string"
	"github.com/ruriazz/gopen-api/src/models"
)

func (uc CorsUsecase) Hostname() domainInterface.HostnameUsecases {
	usecase := HostnameUsecase{&uc}

	usecase.Logger = logger.NewExecutionLog(constants.USECASE_MODULE, "corsUsecase", "HostnameUsecase")
	return usecase
}

func (uc HostnameUsecase) RegisterV1(registerData domainEntity.RegisterDataV1) (*models.Consumer, error) {
	// captcha, err := captchaHelper.NewHCaptcha(uc.Manager)
	// if err != nil {
	// 	return nil, err
	// }

	// if err := captcha.ResponseValidation(registerData.ResponseToken); err != nil {
	// 	return nil, errors.New("E0006")
	// }

	consumer, err := uc.Repositories.Hostname().SingleConsumerV1(models.Consumer{Hostname: registerData.Hostname})
	if err != nil {
		return nil, err
	}

	if consumer != nil {
		return nil, errors.New("E1000")
	}

	secretKey, err := stringHelper.RandomStringURLSafe(32)
	if err != nil {
		return nil, errors.New("E1001")
	}

	hiddenSecretKey, err := encryptionHelper.CreatePassword(secretKey)
	if err != nil {
		return nil, errors.New("E1001")
	}

	consumer, err = uc.Repositories.Hostname().CreateOneV1(models.Consumer{
		Hostname:       registerData.Hostname,
		MaintenerEmail: registerData.Email,
		IsActive:       true,
		SecretKey:      hiddenSecretKey,
	})

	if err != nil {
		return nil, errors.New("E1001")
	}

	consumer.SecretKey = secretKey
	return consumer, nil
}

func (uc HostnameUsecase) CreateChallengeV1(consumer models.Consumer, challengeData domainEntity.CreateChallengeV1) (*models.ConsumerChallenge, error) {
	// captcha, err := captchaHelper.NewHCaptcha(uc.Manager)
	// if err != nil {
	// 	return nil, err
	// }

	// if err := captcha.ResponseValidation(challengeData.ResponseToken); err != nil {
	// 	return nil, errors.New("E0006")
	// }

	if consumer.IsValidated {
		return nil, errors.New("E1002")
	}

	consumerChallenge := models.ConsumerChallenge{ConsumerPkid: consumer.Pkid}
	if err := uc.Repositories.Hostname().GetOneConsumerChallengeV1(&consumerChallenge, false); err != nil {
		if errors.Is(err, constants.ErrorRecordNotFound) {
			consumerChallenge = models.ConsumerChallenge{}
		} else {
			return nil, errors.New("E1003")
		}
	}

	challengeValue, err := stringHelper.RandomStringURLSafe(24)
	if err != nil {
		return nil, errors.New("E1003")
	}

	hiddenChallengeValue := ""
	switch challengeData.ChallengeType {
	case constants.ACME.String():
		hiddenChallengeValue, err = encryptionHelper.StringEnrypt(challengeValue, constants.AES_CBC)
		if err != nil {
			return nil, errors.New("E1003")
		}
	case constants.DNS.String():
		challengeValue = fmt.Sprintf("__%s.%s", consumer.Hostname, challengeValue)
		hiddenChallengeValue, err = encryptionHelper.CreatePassword(challengeValue)
		if err != nil {
			return nil, errors.New("E1003")
		}
	}

	if (consumerChallenge == models.ConsumerChallenge{}) {
		consumerChallenge = models.ConsumerChallenge{
			IsActive:       true,
			ConsumerPkid:   consumer.Pkid,
			ChallengeType:  challengeData.ChallengeType,
			ChallengeValue: hiddenChallengeValue,
			ExpiredAt:      time.Now().Add(24 * time.Hour),
		}

		if err := uc.Repositories.Hostname().CeateConsumerChallengeV1(&consumerChallenge); err != nil {
			return nil, errors.New("E1003")
		}
	} else {
		if consumerChallenge.ExpiredAt.After(time.Now()) {
			return nil, errors.New("E1004")
		}

		newData := models.ConsumerChallenge{
			IsActive:       true,
			ChallengeType:  challengeData.ChallengeType,
			ChallengeValue: hiddenChallengeValue,
			ExpiredAt:      time.Now().Add(24 * time.Hour),
		}

		if err := uc.Repositories.Hostname().UpdateConsumerChallengeV1(&consumerChallenge, newData); err != nil {
			return nil, errors.New("E1003")
		}
	}

	consumerChallenge.ChallengeValue = challengeValue
	consumerChallenge.Consumer = consumer

	// TODO: Send challenge info to maintener email

	return &consumerChallenge, nil
}
