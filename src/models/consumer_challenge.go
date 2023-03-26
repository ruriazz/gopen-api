package models

import "time"

type ConsumerChallenge struct {
	Pkid           int64 `gorm:"primaryKey" json:"-"`
	ConsumerPkid   int64 `gorm:"index" json:"-"`
	IsActive       bool
	IsVerified     bool
	ChallengeType  string
	ChallengeValue string
	ExpiredAt      time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time `gorm:"index;default:NULL" json:"-"`

	Consumer Consumer `gorm:"foreignKey:ConsumerPkid;references:Pkid"`
}

func (ConsumerChallenge) TableName() string {
	return "consumer_challenge"
}
