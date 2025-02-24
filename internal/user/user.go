package user

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/itsjustabavaar/oh-my-gosh/internal/database"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GenerateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

func HashPassword(password, salt string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func AddUser(username, password string) (*User, error) {
	var existingUser User
	if err := database.GetDB().Where("username = ?", username).First(&existingUser).Error; err == nil {
		return nil, utils.ErrUserAlreadyExists
	}
	salt, err := GenerateSalt()
	if err != nil {
		return nil, err
	}
	hashedPassword, err := HashPassword(password, salt)
	if err != nil {
		return nil, err
	}
	user := &User{
		Username:     username,
		PasswordHash: hashedPassword,
		Salt:         salt,
		LastLogin:    time.Now(),
	}
	if err := database.GetDB().Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func Login(username, password string) (*User, error) {
	var user User
	if err := database.GetDB().Where("username = ?", username).First(&user).Error; err != nil {
		return nil, utils.ErrUserNotFound
	}
	hashedPassword, err := HashPassword(password, user.Salt)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(hashedPassword)); err != nil {
		return nil, utils.ErrInvalidPassword
	}
	user.LastLogin = time.Now()
	if err := database.GetDB().Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
