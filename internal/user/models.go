package user

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	LastLogin time.Time
}
