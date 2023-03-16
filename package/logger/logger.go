package logger

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func NewLogger() Log {
	opt := new(LogOptions)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	return opt
}

func (lo *LogOptions) ApiLog(c *gin.Context, t time.Time) {

	url, err := url.Parse(fmt.Sprintf("http://%s", c.Request.Host))
	if err != nil {
		fmt.Println("error", err)
	}

	param := make(map[string]interface{}, len(c.Request.URL.Query()))
	for k, v := range c.Request.URL.Query() {
		param[k] = v[0]
	}
	if len(param) == 0 {
		param = nil
	}

	logResponse := log.WithFields(log.Fields{
		"attributes": &ResponseLog{
			Hostname:       strings.TrimPrefix(url.Hostname(), "www."),
			Fullpath:       c.FullPath(),
			Method:         c.Request.Method,
			Parameters:     param,
			Headers:        c.Request.Header,
			UserAgent:      c.Request.UserAgent(),
			Latency:        time.Since(t).String(),
			ResponseStatus: c.Writer.Status(),
		},
	})

	logResponse.Info("HTTP Call Trace")
}
