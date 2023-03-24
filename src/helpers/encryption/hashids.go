package encryptionHelper

import (
	"fmt"

	"github.com/ruriazz/gopen-api/package/settings"
	"github.com/speps/go-hashids/v2"
)

func getHashids(salt string) (*hashids.HashID, error) {
	settings, err := settings.NewSettings()
	if err != nil {
		return nil, err
	}

	salt = fmt.Sprintf("%s@%s", settings.HASHIDS_SALT, salt)
	data := hashids.NewData()
	data.Salt = salt
	data.MinLength = 12

	return hashids.NewWithData(data)
}

func EncodeID(salt string, id int64) (string, error) {
	encoder, err := getHashids(salt)
	if err != nil {
		return "", err
	}

	encoded, err := encoder.EncodeInt64([]int64{id})
	if err != nil {
		return "", nil
	}
	return encoded, nil
}

func DecodeHashid(salt string, hash string) (int64, error) {
	decoder, err := getHashids(salt)
	if err != nil {
		return 0, err
	}

	decoded, err := decoder.DecodeInt64WithError(hash)
	if err != nil {
		return 0, err
	}

	return decoded[0], nil
}
