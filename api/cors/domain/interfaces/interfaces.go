package corsDomainInterface

type CorsHandlers interface {
	Hostname() HostnameHandlers
}

type CorsUsecases interface {
	Hostname() HostnameUsecases
}

type CorsRepositories interface {
	Hostname() HostnameRepositories
}

type CorsSerializers interface {
	Hostname() HostnameSerializers
}

type CorsValidators interface {
	Hostname() HostnameValidators
}
