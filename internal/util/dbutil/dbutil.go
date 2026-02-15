package dbutil

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db      *gorm.DB
	dbError error
)

func init() {
	db, dbError = gorm.Open(sqlite.Open("gosh.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}

func GormDB() (*gorm.DB, error) {
	return db, dbError
}
