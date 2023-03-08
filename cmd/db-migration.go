package main

import (
	"fmt"
	"os"
	"time"

	goose "github.com/pressly/goose/v3"
	"github.com/ruriazz/gopen-api/package/config"
	"github.com/ruriazz/gopen-api/package/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type gooseDbVersion struct {
	ID        uint      `gorm:"primaryKey,column:id"`
	VersionID uint      `gorm:"column:version_id"`
	IsApplied bool      `gorm:"column:is_applied"`
	TStamp    time.Time `gorm:"column:tstamp"`
}

func (gooseDbVersion) TableName() string {
	return "goose_db_version"
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		args = append(args, "status")
	}

	config, err := config.CreateConfig()
	if err != nil {
		panic(err)
	}

	database, err := database.NewDatabase(config)
	if err != nil {
		panic(err)
	}

	_db, err := database.MySQLGorm.DB()
	if err != nil {
		panic(err)
	}

	if config.APP_ENV != "production" {
		fmt.Printf("Connected to: \"%s\"\n", config.MYSQL_DSN)
	}

	goose.EnsureDBVersion(_db)

	if err := ensureZeroVersion(database.MySQLGorm); err != nil {
		panic(err)
	}

	options := []goose.OptionsFunc{
		goose.WithAllowMissing(),
	}

	if err := goose.RunWithOptions(args[0], _db, "./migrations", []string{}, options...); err != nil {
		panic(err)
	}
}

func ensureZeroVersion(db *gorm.DB) error {
	zeroVersion := gooseDbVersion{ID: 1, VersionID: 0, IsApplied: true, TStamp: time.Now()}
	result := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&zeroVersion)
	return result.Error
}
