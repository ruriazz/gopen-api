package corsUsecase

import (
	"errors"

	domainEntity "github.com/ruriazz/gopen-api/api/cors/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"

	// captchaHelper "github.com/ruriazz/gopen-api/src/helpers/captcha"
	encryptionHelper "github.com/ruriazz/gopen-api/src/helpers/encryption"
	stringHelper "github.com/ruriazz/gopen-api/src/helpers/string"
	"github.com/ruriazz/gopen-api/src/models"
)

func (uc CorsUsecase) Hostname() domainInterface.HostnameUsecases {
	return HostnameUsecase{&uc}
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
		return nil, err
	}

	hiddenSecretKey, err := encryptionHelper.CreatePassword(secretKey)
	if err != nil {
		return nil, err
	}

	consumer, err = uc.Repositories.Hostname().CreateOneV1(models.Consumer{
		Hostname:       registerData.Hostname,
		MaintenerEmail: registerData.Email,
		IsActive:       true,
		SecretKey:      hiddenSecretKey,
	})
	if err != nil {
		return nil, err
	}

	consumer.SecretKey = secretKey
	return consumer, nil
}
