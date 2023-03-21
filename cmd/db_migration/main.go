package main

import (
	"fmt"
	"os"
	"time"

	goose "github.com/pressly/goose/v3"
	"github.com/ruriazz/gopen-api/package/databases"
	"github.com/ruriazz/gopen-api/package/settings"
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

	setting, err := settings.NewSettings()
	if err != nil {
		panic(err)
	}

	database, err := databases.NewDatabases(setting).ConnectMySql()
	if err != nil {
		panic(err)
	}

	_db, err := database.DB()
	if err != nil {
		panic(err)
	}

	if setting.APP_ENV != "production" {
		fmt.Printf("Connected to: \"%s\"\n", setting.MYSQL_DSN)
	}

	goose.SetDialect(database.Config.Name())
	goose.EnsureDBVersion(_db)

	if err := ensureZeroVersion(database); err != nil {
		panic(err)
	}

	options := []goose.OptionsFunc{
		goose.WithAllowMissing(),
	}

	if args[0] != "status" {
		if err := goose.Status(_db, "src/migrations"); err != nil {
			panic(err)
		}
	}

	if err := goose.RunWithOptions(args[0], _db, "src/migrations", []string{}, options...); err != nil {
		panic(err)
	}
}

func ensureZeroVersion(db *gorm.DB) error {
	zeroVersion := gooseDbVersion{ID: 1, VersionID: 0, IsApplied: true, TStamp: time.Now()}
	result := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&zeroVersion)
	return result.Error
}
