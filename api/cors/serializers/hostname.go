package corsSerializer

import domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"

func (s CorsSerializer) Hostname() domainInterface.HostnameSerializers {
	return HostnameSerializer{&s}
}
