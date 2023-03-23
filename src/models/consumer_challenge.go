package models

type ConsumerChallenge struct {
	Pkid         int `gorm:"primaryKey" json:"-"`
	ConsumerPkid int `gorm:"index"`
}

func (ConsumerChallenge) TableName() string {
	return "consumer_challenge"
}
