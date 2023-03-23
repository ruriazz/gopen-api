package encryptionHelper

import (
	"github.com/ruriazz/gopen-api/package/settings"
	"github.com/ruriazz/gopen-api/src/constants"
	"golang.org/x/crypto/bcrypt"
)

func getSecretKey() ([]byte, error) {
	settings, err := settings.NewSettings()
	if err != nil {
		return nil, err
	}

	return []byte(settings.SECRET_KEY), nil
}

func StringEnrypt(plainText string, mode constants.EncryptioMode) (string, error) {
	secretKey, err := getSecretKey()
	if err != nil {
		return "", err
	}

	switch mode {
	case constants.AES_CBC:
		return encryptAESCBC(secretKey, []byte(plainText))
	case constants.AES_GCM:
		return encryptAESGCM(secretKey, []byte(plainText))
	default:
		return "", nil
	}
}

func StringDecrypt(hash string, mode constants.EncryptioMode) ([]byte, error) {
	secretKey, err := getSecretKey()
	if err != nil {
		return nil, err
	}

	switch mode {
	case constants.AES_CBC:
		return decryptAESCBC(secretKey, hash)
	case constants.AES_GCM:
		return decryptAESGCM(secretKey, hash)
	default:
		return nil, nil
	}
}

func CreatePassword(plainText string) (string, error) {
	secretKey, err := getSecretKey()
	if err != nil {
		return "", err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return encryptAESCBC(secretKey, password)
}

func PasswordValidation(plainPassword string, hashPassword string) (bool, error) {
	secretKey, err := getSecretKey()
	if err != nil {
		return false, err
	}

	password, err := decryptAESCBC(secretKey, hashPassword)
	if err != nil {
		return false, err
	}

	if err = bcrypt.CompareHashAndPassword(password, []byte(plainPassword)); err != nil {
		return false, err
	}

	return true, nil
}
