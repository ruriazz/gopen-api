package database

import "gorm.io/gorm"

type Databases struct {
	MySQLGorm *gorm.DB
}
