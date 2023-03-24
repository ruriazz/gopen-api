package captchaHelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ruriazz/gopen-api/package/manager"
	"github.com/ruriazz/gopen-api/package/settings"
	sliceHelper "github.com/ruriazz/gopen-api/src/helpers/slice"
)

type hCaptcha struct {
	settings    hCaptchaSettings
	appSettings settings.Setting
}

type HCaptcha interface {
	GetSecretKey() string
	GetApiURI() string
	ResponseValidation(token string) error
}

func NewHCaptcha(manager manager.Manager) (HCaptcha, error) {
	settings, err := getSettings(manager)
	if err != nil {
		return nil, err
	}

	return hCaptcha{
		settings:    *settings,
		appSettings: *manager.Settings,
	}, nil
}

func getSettings(manager manager.Manager) (*hCaptchaSettings, error) {
	return &hCaptchaSettings{
		secretKey: manager.Settings.HCAPTCHA_SECRET_KEY,
		apiURI:    manager.Settings.HCAPTCHA_API_URI,
		ctsLimit:  manager.Settings.HCAPTCHA_CTS_LIMIT,
	}, nil
}

func (c hCaptcha) GetSecretKey() string {
	return c.settings.secretKey
}

func (c hCaptcha) GetApiURI() string {
	return c.settings.apiURI
}

func (c hCaptcha) ResponseValidation(token string) error {
	var data apiResponse
	var client = &http.Client{}
	var param = url.Values{}

	param.Set("secret", c.GetSecretKey())
	param.Set("response", token)

	request, err := http.NewRequest("POST", c.GetApiURI(), bytes.NewBufferString(param.Encode()))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return err
	}

	if !data.Success {
		return fmt.Errorf("(%s) API Response Error: [%s]", c.GetApiURI(), strings.Join(data.ErrorCodes[:], ", "))
	}

	if time.Since(data.ChallengeTS).Minutes() >= c.settings.ctsLimit {
		return fmt.Errorf("(%s) challenge time stamp limit", c.GetApiURI())
	}

	if !sliceHelper.StringInSlice(data.Hostname, c.appSettings.HTTP_ALLOWED_HOSTS) {
		return fmt.Errorf("(%s) %s is not allowed hostname", c.GetApiURI(), data.Hostname)
	}

	return nil
}
