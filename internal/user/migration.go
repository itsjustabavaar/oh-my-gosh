package user

import (
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}
