package authentication

import (
	"errors"

	"github.com/ruriazz/gopen-api/package/databases"
	"github.com/ruriazz/gopen-api/package/logger"
	"github.com/ruriazz/gopen-api/src/constants"
	"github.com/ruriazz/gopen-api/src/models"
	"gorm.io/gorm"
)

type authenticationRepository struct {
	Database databases.Database
	Logger   logger.ExecutionLog
}

type authenticationRepositoryInterface interface {
	GetOneConsumer(model *models.Consumer) error
}

func newAuthenticationRepository(database databases.Database) authenticationRepositoryInterface {
	return authenticationRepository{
		Database: database,
		Logger:   logger.NewExecutionLog(constants.PACKAGE_MODULE, "authentication", "authenticationRepository"),
	}
}

func (r authenticationRepository) GetOneConsumer(model *models.Consumer) error {
	err := r.Database.MySqlDB.
		Model(model).
		Find(&model).
		First(&model).
		Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			r.Logger.Error("GetOneConsumer", 1, err.Error(), nil)
		}

		return err
	}

	return nil
}
