package restserver

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ruriazz/gopen-api/package/settings"
)

func NewRestServer(conf *settings.Setting) (*Server, error) {
	if conf.APP_ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	_engine := gin.New()
	_http := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", conf.HTTP_SERVER_HOST, conf.HTTP_SERVER_PORT),
		Handler:      _engine,
		ReadTimeout:  time.Duration(conf.HTTP_READ_TIMEOUT) * time.Second,
		WriteTimeout: time.Duration(conf.HTTP_WRITE_TIMEOUT) * time.Second,
	}

	_server := &Server{
		Engine:     _engine,
		HttpServer: _http,
	}

	return _server, nil
}
