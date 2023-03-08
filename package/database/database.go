package database

import "github.com/ruriazz/gopen-api/package/config"

func NewDatabase(config *config.Config) (*Databases, error) {
	var db Databases

	mysql, err := db.MySQLConnect(config.MYSQL_DSN)
	if err != nil {
		return nil, err
	}

	return &Databases{
		MySQLGorm: mysql,
	}, nil
}
