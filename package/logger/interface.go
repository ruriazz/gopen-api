package logger

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Log interface {
	// ErrorLog(ctx gin.Context, err error, addInfo interface{})
	// WarnLog(ctx gin.Context, err error, addInfo interface{})
	// CustomLog(r http.Request, level string, data interface{})
	ApiLog(c *gin.Context, t time.Time)
}

type LogOptions struct{}
