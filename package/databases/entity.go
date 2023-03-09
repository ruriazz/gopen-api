package databases

import "gorm.io/gorm"

type Database struct {
	MySqlDB *gorm.DB
	RedisDB interface{}
}

type Options struct {
	mySqlDSN string
}
