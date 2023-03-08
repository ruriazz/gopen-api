package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ruriazz/gopen-api/package/config"
)

func CreateCors(config *config.Config) gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = config.CORS_ALLOWED_HEADERS

	if config.APP_ENV == "production" {
		corsConfig.AllowAllOrigins = false
		corsConfig.AllowOrigins = config.CORS_ALLOWED_ORIGINS
		corsConfig.AllowMethods = config.CORS_ALLOWED_METHODS
	} else {
		corsConfig.AllowAllOrigins = true
	}

	return cors.New(corsConfig)
}
