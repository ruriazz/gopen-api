package manager

import (
	"github.com/ruriazz/gopen-api/package/authentication"
	"github.com/ruriazz/gopen-api/package/databases"
	"github.com/ruriazz/gopen-api/package/logger"
	restserver "github.com/ruriazz/gopen-api/package/rest_server"
	"github.com/ruriazz/gopen-api/package/settings"
)

type Manager struct {
	Settings       *settings.Setting
	Server         *restserver.Server
	Databases      *databases.Database
	Logger         *logger.LogOptions
	Authentication authentication.Authentication
}
