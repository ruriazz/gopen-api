package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ruriazz/gopen-api/package/settings"
)

func newCors(settings *settings.Setting) gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = settings.CORS_ALLOWED_HEADERS

	if settings.APP_ENV == "production" {
		corsConfig.AllowAllOrigins = false
		corsConfig.AllowOrigins = settings.CORS_ALLOWED_ORIGINS
		corsConfig.AllowMethods = settings.CORS_ALLOWED_METHODS
	} else {
		corsConfig.AllowAllOrigins = true
	}

	return cors.New(corsConfig)
}
