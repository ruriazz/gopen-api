package corsRepository

import (
	domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"
	"github.com/ruriazz/gopen-api/src/models"
)

func (r CorsRepository) Hostname() domainInterface.HostnameRepositories {
	return HostnameRepository{&r}
}

func (r HostnameRepository) SingleConsumerV1(model models.Consumer) (*models.Consumer, error) {
	err := r.Databases.MySqlDB.
		Where(&model).
		First(&model).
		Error

	if err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}

		return nil, err
	}

	return &model, nil
}

func (r HostnameRepository) CreateOneV1(model models.Consumer) (*models.Consumer, error) {
	if err := r.Databases.MySqlDB.Create(&model).Error; err != nil {
		return nil, err
	}

	return &model, nil
}
