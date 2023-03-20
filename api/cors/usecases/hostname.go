package corsUsecase

import domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"

func (uc CorsUsecase) Hostname() domainInterface.HostnameUsecases {
	return HostnameUsecase{&uc}
}
