package corsValidator

import domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"

func (v CorsValidator) Hostname() domainInterface.HostnameValidators {
	return HostnameValidator{&v}
}
