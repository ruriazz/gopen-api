package models

type Consumer struct {
	Pkid         int16 `gorm:"primaryKey" json:"-"`
	Hostname     string
	EmailAddress string
	PrivateKey   string
}

func (Consumer) TableName() string {
	return "consumer"
}
