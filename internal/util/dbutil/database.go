package dbutil

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	PostgresHost     = "0.0.0.0"
	PostgresPort     = "5432"
	PostgresUser     = "myuser"
	PostgresPassword = "mypassword"
	PostgresDb       = "mydb"
)

var (
	db      *gorm.DB
	dbError error
)

func init() {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		PostgresHost,
		PostgresPort,
		PostgresDb,
		PostgresUser,
		PostgresPassword)

	db, dbError = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

func GormDB() (*gorm.DB, error) {
	return db, dbError
}
