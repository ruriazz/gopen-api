package corsDomainInterface

import "github.com/gin-gonic/gin"

type HostnameHandlers interface {
	GetInfoV1(context *gin.Context)
	RegisterV1(context *gin.Context)
	NewChallenge(context *gin.Context)
}

type HostnameUsecases interface{}

type HostnameRepositories interface{}

type HostnameSerializers interface{}

type HostnameValidators interface{}
