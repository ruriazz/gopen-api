package corsValidator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	domainEntity "github.com/ruriazz/gopen-api/api/cors/domain/entities"
	domainInterface "github.com/ruriazz/gopen-api/api/cors/domain/interfaces"
	"github.com/ruriazz/gopen-api/src/constants"
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

func (v HostnameValidator) CreateChallengeV1(context *gin.Context) error {
	var data domainEntity.CreateChallengeV1
	if err := context.ShouldBindJSON(&data); err != nil {
		return err
	}

	if data.ChallengeType == "" || data.ResponseToken == "" {
		return errors.New("(HostnameValidator.CreateChallengeV1) invalid required fields")
	}

	if data.ChallengeType != constants.ACME.String() && data.ChallengeType != constants.DNS.String() {
		return fmt.Errorf("(HostnameValidator.CreateChallengeV1) '%s' is invalid type", data.ChallengeType)
	}

	context.Set("data", data)
	return nil
}
