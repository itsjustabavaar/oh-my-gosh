package user

import (
	"github.com/itsjustabavaar/oh-my-gosh/internal/database"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

func VerifyPassword(inputPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))

	return err == nil
}

func AddUser(username, password string) error {
	var existingUser User

	if err := database.GetDB().Where("username = ?", username).First(&existingUser).Error; err == nil {
		return utils.ErrUserAlreadyExists
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	user := &User{
		Username:  username,
		Password:  hashedPassword,
		LastLogin: time.Now(),
	}

	if err = database.GetDB().Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func Login(username, password string) (*User, error) {
	var user User

	if err := database.GetDB().Where("username = ?", username).First(&user).Error; err != nil {
		return nil, utils.ErrUserNotFound
	}

	ok := VerifyPassword(password, user.Password)
	if !ok {
		return nil, utils.ErrInvalidPassword
	}

	user.LastLogin = time.Now()

	if err := database.GetDB().Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
