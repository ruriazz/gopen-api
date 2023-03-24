package captchaHelper

import "time"

type hCaptchaSettings struct {
	secretKey string
	apiURI    string
	ctsLimit  float64
}

type apiResponse struct {
	Success     bool      `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}
