package manager

import (
	"github.com/ruriazz/gopen-api/package/config"
	"github.com/ruriazz/gopen-api/package/database"
	"github.com/ruriazz/gopen-api/package/middleware"
	responseWritter "github.com/ruriazz/gopen-api/package/response-writter"
	"github.com/ruriazz/gopen-api/package/server"
)

func CreateManager() (*Manager, error) {
	_config, err := config.CreateConfig()
	if err != nil {
		return nil, err
	}

	_database, err := database.NewDatabase(_config)
	if err != nil {
		return nil, err
	}

	_server, err := server.CreateServer(_config)
	if err != nil {
		return nil, err
	}

	_rw, err := responseWritter.CreateResponseWritter(_config)
	if err != nil {
		return nil, err
	}

	middleware.InitMiddleware(_config, _server)

	return &Manager{
		Conf:      _config,
		Server:    _server,
		RW:        _rw,
		Databases: _database,
	}, nil
}
