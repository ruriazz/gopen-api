package corsDomainEntity

type RegisterDataV1 struct {
	Hostname      string `json:"hostname"`
	Email         string `json:"email"`
	ResponseToken string `json:"responseToken"`
}

type DefaultConsumerInfoV1 struct {
	Hostname       string `json:"hostname"`
	MaintenerEmail string `json:"maintenerEmail"`
	IsActive       bool   `json:"isActive"`
	IsValidated    bool   `json:"isValidated"`
	SecretKey      string `json:"secretKey"`
}
