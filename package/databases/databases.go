package databases

import (
	"time"

	"github.com/cenkalti/backoff"
	"github.com/ruriazz/gopen-api/package/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabases(settings *settings.Setting) Databases {
	opt := new(Options)
	opt.mySqlDSN = settings.MYSQL_DSN

	return opt
}

func (o Options) ConnectMySql() (*gorm.DB, error) {
	gormdb, err := gorm.Open(mysql.Open(o.mySqlDSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := gormdb.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(1000)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := backoff.Retry(func() error {
		if err := sqlDB.Ping(); err != nil {
			return err
		}

		return nil
	}, backoff.NewExponentialBackOff()); err != nil {
		return nil, err
	}

	return gormdb, nil
}

func (o Options) ConnectRedis() (interface{}, error) {
	return nil, nil
}
