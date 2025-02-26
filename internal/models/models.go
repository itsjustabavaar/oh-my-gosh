package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	LastLogin time.Time
	History   []CommandHistory `gorm:"foreignKey:UserID"`
}

type CommandHistory struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Command   string    `gorm:"not null"`
	Timestamp time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UserID    uint      `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID"`
}
