package encryptionHelper

import "golang.org/x/crypto/bcrypt"

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
