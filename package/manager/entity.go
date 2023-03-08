package manager

import (
	"github.com/ruriazz/gopen-api/package/config"
	"github.com/ruriazz/gopen-api/package/database"
	"github.com/ruriazz/gopen-api/package/logger"
	responseWritter "github.com/ruriazz/gopen-api/package/response-writter"
	"github.com/ruriazz/gopen-api/package/server"
)

type Manager struct {
	Conf      *config.Config
	Server    *server.Server
	Logger    *logger.LogOptions
	RW        *responseWritter.ResponseWritter
	Databases *database.Databases
}
