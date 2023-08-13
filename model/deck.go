package model

import (
	"time"

	"gorm.io/gorm"
)

type Deck struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:255;not null" json:"name"`
	CreatedAt   time.Time `gorm:"not null" json:"-"`
	UpdatedAt   time.Time `gorm:"not null" json:"-"`
	NmUserDecks []NmUserDeck
	Cards       []Card
}
