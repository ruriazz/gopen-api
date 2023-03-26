package corsDomainEntity

import "time"

type RegisterDataV1 struct {
	Hostname      string `json:"hostname"`
	Email         string `json:"email"`
	ResponseToken string `json:"responseToken"`
}

type CreateChallengeV1 struct {
	ChallengeType string `json:"challengeType"`
	ResponseToken string `json:"responseToken"`
}

type DefaultConsumerInfoV1 struct {
	Hostname       string `json:"hostname"`
	MaintenerEmail string `json:"maintenerEmail"`
	IsActive       bool   `json:"isActive"`
	IsValidated    bool   `json:"isValidated"`
	SecretKey      string `json:"secretKey"`
}

type RecordData struct {
	Target string `json:"target"`
	Type   string `json:"type"`
	Data   string `json:"data"`
}
type DefaultCreateDNSChallengeV1 struct {
	ChallengeType string     `json:"challengeType"`
	RecordData    RecordData `json:"recordData"`
	ExpiredAt     time.Time  `json:"expiredAt"`
}

type DefaultCreateACMEChallengeV1 struct {
	ChallengeType string    `json:"challengeType"`
	ACMEFile      string    `json:"acmeFile"`
	ExpiredAt     time.Time `json:"expiredAt"`
}
