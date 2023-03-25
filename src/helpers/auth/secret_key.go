package authHelper

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	corsRepository "github.com/ruriazz/gopen-api/api/cors/repositories"
	"github.com/ruriazz/gopen-api/package/manager"
	"github.com/ruriazz/gopen-api/src/constants"
	encryptionHelper "github.com/ruriazz/gopen-api/src/helpers/encryption"
	responseHelper "github.com/ruriazz/gopen-api/src/helpers/response"
	"github.com/ruriazz/gopen-api/src/models"
)

func SecretKeyAuth(manager manager.Manager) gin.HandlerFunc {
	repo := corsRepository.NewCorsRepository(manager)
	return func(context *gin.Context) {
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

		var data models.Consumer
		if err := data.SetPkid(secrets[0], &data); err != nil {
			responseHelper.JSON(responseHelper.FieldsV1{
				Context: context,
				Error:   errors.New("E0005"),
			})
			return
		}

		consumer, err := repo.Hostname().SingleConsumerV1(data)
		if err != nil || consumer == nil {
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
		context.Set(constants.VAR_CONSUMER_DATA, *consumer)
		context.Next()
	}
}
