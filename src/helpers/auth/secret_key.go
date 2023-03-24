package authHelper

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ruriazz/gopen-api/package/manager"
	"github.com/ruriazz/gopen-api/src/constants"
	encryptionHelper "github.com/ruriazz/gopen-api/src/helpers/encryption"
	responseHelper "github.com/ruriazz/gopen-api/src/helpers/response"
	"github.com/ruriazz/gopen-api/src/models"
)

func SecretKeyAuth(manager manager.Manager) gin.HandlerFunc {
	db := manager.Databases.MySqlDB
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

		if err := db.Model(&data).First(&data).Error; err != nil {
			responseHelper.JSON(responseHelper.FieldsV1{
				Context: context,
				Error:   errors.New("E0005"),
			})
			return
		}

		if valid, _ := encryptionHelper.PasswordValidation(secrets[1], data.SecretKey); !valid {
			responseHelper.JSON(responseHelper.FieldsV1{
				Context: context,
				Error:   errors.New("E0005"),
			})
			return
		}

		data.SecretKey = secrets[1]
		context.Set(constants.VAR_CONSUMER_DATA, data)
		context.Next()
	}
}
