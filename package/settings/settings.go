package settings

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

func NewSettings() (*Setting, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.SetConfigName("master.env")

			if err := viper.ReadInConfig(); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	_config := &Setting{}
	if err := viper.Unmarshal(&_config); err != nil {
		return nil, err
	}

	_config.CORS_ALLOWED_ORIGINS = []string{}
	if viper.Get("CORS_ALLOWED_ORIGINS") != nil {
		for _, s := range strings.Split(viper.GetString("CORS_ALLOWED_ORIGINS"), ",") {
			_config.CORS_ALLOWED_ORIGINS = append(_config.CORS_ALLOWED_ORIGINS, strings.TrimSpace(s))
		}
	}

	_config.CORS_ALLOWED_HEADERS = []string{}
	if viper.Get("CORS_ALLOWED_HEADERS") != nil {
		for _, s := range strings.Split(viper.GetString("CORS_ALLOWED_HEADERS"), ",") {
			_config.CORS_ALLOWED_HEADERS = append(_config.CORS_ALLOWED_HEADERS, strings.TrimSpace(s))
		}
	}

	_config.CORS_ALLOWED_METHODS = []string{}
	if viper.Get("CORS_ALLOWED_METHODS") != nil {
		for _, s := range strings.Split(viper.GetString("CORS_ALLOWED_METHODS"), ",") {
			_config.CORS_ALLOWED_METHODS = append(_config.CORS_ALLOWED_METHODS, strings.TrimSpace(s))
		}
	}

	_config.HTTP_ALLOWED_HOSTS = []string{}
	if viper.Get("HTTP_ALLOWED_HOSTS") != nil {
		for _, s := range strings.Split(viper.GetString("HTTP_ALLOWED_HOSTS"), ",") {
			_config.HTTP_ALLOWED_HOSTS = append(_config.HTTP_ALLOWED_HOSTS, strings.TrimSpace(s))
		}
	}

	return _config, nil
}
