package corsValidator

import (
	"errors"

	"github.com/gin-gonic/gin"
	domainEntity "github.com/ruriazz/gopen-api/api/cors/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"
)

type HostnameValidator struct{ *CorsValidator }

func (v CorsValidator) Hostname() domainInterface.HostnameValidators {
	return HostnameValidator{&v}
}

func (v HostnameValidator) RegisterV1(context *gin.Context) (*gin.Context, error) {
	var data domainEntity.RegisterDataV1
	if err := context.ShouldBindJSON(&data); err != nil {
		return context, err
	}

	if data.Hostname == "" || data.Email == "" || data.ResponseToken == "" {
		return context, errors.New("")
	}

	context.Set("data", data)
	return context, nil
}
