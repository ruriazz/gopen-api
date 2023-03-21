package corsDomainEntity

type RegisterDataV1 struct {
	Hostname      string `json:"hostname"`
	Email         string `json:"email"`
	ResponseToken string `json:"responseToken"`
}
