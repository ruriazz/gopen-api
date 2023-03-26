package models

import (
	"errors"
	"time"

	encryptionHelper "github.com/ruriazz/gopen-api/src/helpers/encryption"
)

type Consumer struct {
	Pkid           int64  `gorm:"primaryKey" json:"-"`
	Hostname       string `gorm:"index"`
	MaintenerEmail string
	IsActive       bool
	IsValidated    bool
	SecretKey      string `json:"-"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time `gorm:"index;default:NULL" json:"-"`
}

func (Consumer) TableName() string {
	return "consumer"
}

func (m Consumer) GetID() string {
	encoded, err := encryptionHelper.EncodeID(m.TableName(), int64(m.Pkid))
	if err != nil {
		return ""
	}

	return encoded
}

func (m Consumer) SetPkid(hash string, model *Consumer) error {
	if hash == "" {
		return errors.New("invalid hash")
	}

	decoded, err := encryptionHelper.DecodeHashid(m.TableName(), hash)
	if err != nil {
		return err
	}

	model.Pkid = decoded
	return nil
}
