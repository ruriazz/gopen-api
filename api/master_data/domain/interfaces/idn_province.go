package masterDataDomainInterface

import "github.com/gin-gonic/gin"

type IdnProvinceHandlers interface {
	GetCollections(ctx *gin.Context)
}

type IdnProvinceUsecases interface{}

type IdnProvinceRepositories interface{}
