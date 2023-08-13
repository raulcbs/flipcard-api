package model

import (
	"gorm.io/gorm"
)

type NmUserDeck struct {
	gorm.Model
	UserID uint
	DeckID uint
}
