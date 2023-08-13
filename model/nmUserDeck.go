package model

import (
	"gorm.io/gorm"
)

type NmUserDeck struct {
	gorm.Model
	ID     uint `gorm:"primaryKey"`
	UserID uint
	DeckID uint
}
