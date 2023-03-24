package settings

type Setting struct {
	APP_ENV              string `mapstructure:"APP_ENV"`
	APP_TZ               string `mapstructure:"APP_TZ"`
	SECRET_KEY           string `mapstructure:"SECRET_KEY"`
	HTTP_SERVER_HOST     string `mapstructure:"HTTP_SERVER_HOST"`
	HTTP_SERVER_PORT     string `mapstructure:"HTTP_SERVER_PORT"`
	HTTP_READ_TIMEOUT    int8   `mapstructure:"HTTP_READ_TIMEOUT"`
	HTTP_WRITE_TIMEOUT   int8   `mapstructure:"HTTP_WRITE_TIMEOUT"`
	HTTP_ALLOWED_HOSTS   []string
	CORS_ALLOWED_ORIGINS []string
	CORS_ALLOWED_HEADERS []string
	CORS_ALLOWED_METHODS []string
	MYSQL_DSN            string  `mapstructure:"MYSQL_DSN"`
	HCAPTCHA_SECRET_KEY  string  `mapstructure:"HCAPTCHA_SECRET_KEY"`
	HCAPTCHA_API_URI     string  `mapstructure:"HCAPTCHA_API_URI"`
	HCAPTCHA_CTS_LIMIT   float64 `mapstructure:"HCAPTCHA_CTS_LIMIT"`
	HASHIDS_SALT         string  `mapstructure:"HASHIDS_SALT"`
}
