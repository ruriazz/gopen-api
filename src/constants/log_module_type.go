package constants

type LogModuleType int

const (
	HANDLER_MODULE LogModuleType = iota
	USECASE_MODULE
	REPOSITORY_MODULE
	VALIDATOR_MODULE
	SERIALIZER_MODULE
	HELPER_MODULE
	PACKAGE_MODULE
)

func (l LogModuleType) String() string {
	return [...]string{"handler", "usecase", "repository", "validator", "serializer", "helper", "package"}[l]
}
