package corsRepository

import (
	"errors"

	domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"
	"github.com/ruriazz/gopen-api/package/logger"
	"github.com/ruriazz/gopen-api/src/constants"
	"github.com/ruriazz/gopen-api/src/models"
	"gorm.io/gorm"
)

func (r CorsRepository) Hostname() domainInterface.HostnameRepositories {
	repository := HostnameRepository{&r}

	repository.Logger = logger.NewExecutionLog(constants.REPOSITORY_MODULE, "corsRepository", "HostnameRepository")
	return repository
}

func (r HostnameRepository) SingleConsumerV1(model models.Consumer) (*models.Consumer, error) {
	err := r.Databases.MySqlDB.
		Where(&model).
		First(&model).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		r.Logger.Error("SingleConsumerV1", 1, err.Error(), nil)
		return nil, err
	}

	return &model, nil
}

func (r HostnameRepository) CreateOneV1(model models.Consumer) (*models.Consumer, error) {
	if err := r.Databases.MySqlDB.Create(&model).Error; err != nil {
		r.Logger.Error("CreateOneV1", 1, err.Error(), nil)
		return nil, err
	}

	return &model, nil
}

func (r HostnameRepository) GetOneConsumerChallengeV1(model *models.ConsumerChallenge, preload bool) error {
	session := r.Databases.MySqlDB.Where(&model)
	if preload {
		session = session.Preload("Consumer")
	}

	if err := session.First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constants.ErrorRecordNotFound
		}

		r.Logger.Error("GetOneConsumerChallengeV1", 1, err.Error(), nil)
		return err
	}
	return nil
}

func (r HostnameRepository) CeateConsumerChallengeV1(model *models.ConsumerChallenge) error {
	if err := r.Databases.MySqlDB.Create(&model).Error; err != nil {
		r.Logger.Error("CeateConsumerChallengeV1", 1, err.Error(), nil)
		return err
	}

	return nil
}

func (r HostnameRepository) UpdateConsumerChallengeV1(model *models.ConsumerChallenge, newData models.ConsumerChallenge) error {
	if err := r.Databases.MySqlDB.Model(&model).Updates(newData).Error; err != nil {
		r.Logger.Error("UpdateConsumerChallengeV1", 1, err.Error(), nil)
		return err
	}

	return nil
}
