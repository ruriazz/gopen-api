package corsValidator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	domainEntity "github.com/ruriazz/gopen-api/api/cors/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"
	stringHelper "github.com/ruriazz/gopen-api/src/helpers/string"
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
		return context, errors.New("(HostnameValidator.RegisterV1) invalid required fields")
	}

	if !stringHelper.IsValidEmail(data.Email) {
		return context, fmt.Errorf("(HostnameValidator.RegisterV1) '%s' is invalid email", data.Email)
	}

	if !stringHelper.IsValidHostname(data.Hostname) {
		return context, fmt.Errorf("(HostnameValidator.RegisterV1) '%s' is invalid hostname", data.Hostname)
	}

	data.Email = strings.ToLower(data.Email)
	context.Set("data", data)
	return context, nil
}
