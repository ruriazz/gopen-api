package corsUsecase

import (
	"fmt"

	domainEntity "github.com/ruriazz/gopen-api/api/cors/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"
	"github.com/ruriazz/gopen-api/src/models"
)

func (uc CorsUsecase) Hostname() domainInterface.HostnameUsecases {
	return HostnameUsecase{&uc}
}

func (uc HostnameUsecase) RegisterV1(registerData domainEntity.RegisterDataV1) (*models.Consumer, error) {
	fmt.Println(registerData)
	return nil, nil
}
