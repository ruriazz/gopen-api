package restserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine     *gin.Engine
	HttpServer *http.Server
}
