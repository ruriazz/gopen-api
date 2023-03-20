package corsRepository

import domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"

func (r CorsRepository) Hostname() domainInterface.HostnameRepositories {
	return HostnameRepository{&r}
}
