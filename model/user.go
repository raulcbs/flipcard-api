package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:255;not null;unique" json:"name"`
	Password    string    `gorm:"size:255;not null" json:"-"`
	CreatedAt   time.Time `gorm:"not null" json:"-"`
	UpdatedAt   time.Time `gorm:"not null" json:"-"`
	NmUserDecks []NmUserDeck
}
