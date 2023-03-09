package databases

import "gorm.io/gorm"

type Databases interface {
	ConnectMySql() (*gorm.DB, error)
	ConnectRedis() (interface{}, error)
}
