package middleware

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ruriazz/gopen-api/package/config"
	"github.com/ruriazz/gopen-api/package/helper"
	"github.com/ruriazz/gopen-api/package/logger"
	"github.com/ruriazz/gopen-api/package/server"
)

func InitMiddleware(config *config.Config, server *server.Server) {
	corsMiddleware := CreateCors(config)

	server.Engine.Use(DefaultMiddleware(config))
	server.Engine.Use(corsMiddleware)
}

func DefaultMiddleware(config *config.Config) gin.HandlerFunc {
	_logger := logger.CreateLogger()
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
		if !helper.StringInSlice(hostname, config.HTTP_ALLOWED_HOSTS) {
			c.String(403, fmt.Sprintf("'%s' is not allowed host", hostname))
			c.Abort()
		}

		c.Set("requestTime", t)
		c.Next()

		_logger.ApiLog(c, t)
	}
}
