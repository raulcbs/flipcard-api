package model

import (
	"gorm.io/gorm"
)

type Deck struct {
	gorm.Model
	Name        string `gorm:"size:255;not null" json:"name"`
	NmUserDecks []NmUserDeck
	Cards       []Card
}
