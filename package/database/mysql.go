package database

import (
	"time"

	"github.com/cenkalti/backoff"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (db *Databases) MySQLConnect(dsn string) (*gorm.DB, error) {
	gormdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
