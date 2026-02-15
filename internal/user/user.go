package user

import (
	"time"

	"github.com/itsjustabavaar/oh-my-gosh/internal/models"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util/dbutil"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util/errorutil"
	"golang.org/x/crypto/bcrypt"
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
	var existingUser models.User

	db, err := dbutil.GormDB()
	if err != nil {
		return err
	}

	if err := db.Where("username = ?", username).First(&existingUser).Error; err == nil {
		return errorutil.ErrUserAlreadyExists
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Username:  username,
		Password:  hashedPassword,
		LastLogin: time.Now(),
	}

	if err = db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(username string) error {
	db, err := dbutil.GormDB()
	if err != nil {
		return err
	}

	if err := db.Where("username = ?", username).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}

func Login(username, password string) (*models.User, error) {
	var user models.User

	db, err := dbutil.GormDB()
	if err != nil {
		return nil, err
	}

	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errorutil.ErrUserNotFound
	}

	ok := VerifyPassword(password, user.Password)
	if !ok {
		return nil, errorutil.ErrInvalidPassword
	}

	user.LastLogin = time.Now()

	if err := db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
