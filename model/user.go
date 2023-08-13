package model

import (
	"html"
	"strings"

	"github.com/raulcbs/flipcard-api/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `gorm:"size:255;not null;unique" json:"name"`
	Password    string `gorm:"size:255;not null" json:"-"`
	NmUserDecks []NmUserDeck
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Name = html.EscapeString(strings.TrimSpace(user.Name))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := database.Database.Where("name=?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserById(id uint) (User, error) {
	var user User
	err := database.Database.Preload("NmUserDecks").Where("ID=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
