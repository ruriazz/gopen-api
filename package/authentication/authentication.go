package authentication

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ruriazz/gopen-api/package/databases"
	"github.com/ruriazz/gopen-api/package/settings"
	"github.com/ruriazz/gopen-api/src/constants"
	encryptionHelper "github.com/ruriazz/gopen-api/src/helpers/encryption"
	responseHelper "github.com/ruriazz/gopen-api/src/helpers/response"
	"github.com/ruriazz/gopen-api/src/models"
)

type authentication struct {
	Setting    settings.Setting
	Database   databases.Database
	Repository authenticationRepositoryInterface
}

type Authentication interface {
	SecretKey(context *gin.Context)
}

func NewAuthentication(setting settings.Setting, database databases.Database) Authentication {
	return authentication{
		Setting:    setting,
		Database:   database,
		Repository: newAuthenticationRepository(database),
	}
}

func (p authentication) SecretKey(context *gin.Context) {
	secret := context.Query("secret")
	if secret == "" {
		responseHelper.JSON(responseHelper.FieldsV1{
			Context: context,
			Error:   errors.New("E0004"),
		})
		return
	}

	secrets := strings.Split(secret, ":")
	if len(secret) < 2 {
		responseHelper.JSON(responseHelper.FieldsV1{
			Context: context,
			Error:   errors.New("E0005"),
		})
		return
	}

	var consumer models.Consumer
	if err := consumer.SetPkid(secrets[0], &consumer); err != nil {
		responseHelper.JSON(responseHelper.FieldsV1{
			Context: context,
			Error:   errors.New("E0005"),
		})
		return
	}

	if err := p.Repository.GetOneConsumer(&consumer); err != nil {
		responseHelper.JSON(responseHelper.FieldsV1{
			Context: context,
			Error:   errors.New("E0005"),
		})
		return
	}

	if valid, _ := encryptionHelper.PasswordValidation(secrets[1], consumer.SecretKey); !valid {
		responseHelper.JSON(responseHelper.FieldsV1{
			Context: context,
			Error:   errors.New("E0005"),
		})
		return
	}

	consumer.SecretKey = secrets[1]
	context.Set(constants.VAR_CONSUMER_DATA, consumer)
	context.Next()
}
