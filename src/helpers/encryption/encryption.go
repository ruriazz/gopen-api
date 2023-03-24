package encryptionHelper

import (
	"github.com/ruriazz/gopen-api/package/settings"
	"github.com/ruriazz/gopen-api/src/constants"
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
