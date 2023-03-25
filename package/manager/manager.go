package manager

import (
	"github.com/ruriazz/gopen-api/package/authentication"
	"github.com/ruriazz/gopen-api/package/databases"
	"github.com/ruriazz/gopen-api/package/middleware"
	restserver "github.com/ruriazz/gopen-api/package/rest_server"
	"github.com/ruriazz/gopen-api/package/settings"
)

func NewManager() (*Manager, error) {
	_settings, err := settings.NewSettings()
	if err != nil {
		return nil, err
	}

	var _db databases.Database
	_databases := databases.NewDatabases(_settings)
	_db.MySqlDB, err = _databases.ConnectMySql()
	if err != nil {
		return nil, err
	}

	_server, err := restserver.NewRestServer(_settings)
	if err != nil {
		return nil, err
	}

	middleware.NewMiddleware(_settings, _server)

	return &Manager{
		Settings:       _settings,
		Databases:      &_db,
		Server:         _server,
		Authentication: authentication.NewAuthentication(*_settings, _db),
	}, nil
}
