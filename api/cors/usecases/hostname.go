package corsUsecase

import (
	"errors"

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
