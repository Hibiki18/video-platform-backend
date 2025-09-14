package models

import (
	"errors"
	"video-platform-backend/config"

	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:50"`
	Password string `gorm:"size:255"`
}

func CreateUser(user *User) error {
	// Hash the password before storing it
	return config.DB.Create(user).Error

}

func GetUserByName(username string) (User, error) {
	var u User
	if err := config.DB.Where("username = ?", username).First(&u).Error; err != nil {
		return u, errors.New("user not found")
	}
	u.Password = "" // Do not return the password
	return u, nil
}