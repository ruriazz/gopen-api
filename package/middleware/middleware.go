package middleware

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ruriazz/gopen-api/package/logger"
	restserver "github.com/ruriazz/gopen-api/package/rest_server"
	"github.com/ruriazz/gopen-api/package/settings"
	helpers "github.com/ruriazz/gopen-api/src/helpers/slice"
)

func NewMiddleware(setting *settings.Setting, server *restserver.Server) {
	corsMiddleware := newCors(setting)

	server.Engine.Use(DefaultMiddleware(setting))
	server.Engine.Use(corsMiddleware)
}

func DefaultMiddleware(settings *settings.Setting) gin.HandlerFunc {
	_logger := logger.NewLogger()
	return func(c *gin.Context) {
		t := time.Now()
		fullpath := c.FullPath()
		if fullpath == "" {
			fullpath = "/"
		}

		url, err := url.Parse(fmt.Sprintf("http://%s", c.Request.Host))
		if err != nil {
			fmt.Println("error", err)
		}

		hostname := strings.TrimPrefix(url.Hostname(), "www.")
		if !helpers.StringInSlice(hostname, settings.HTTP_ALLOWED_HOSTS) {
			c.String(403, fmt.Sprintf("'%s' is not allowed host", hostname))
			c.Abort()
		}

		c = registeredCorsValidation(c)

		c.Set("requestTime", t)
		c.Next()

		_logger.ApiLog(c, t)
	}
}
