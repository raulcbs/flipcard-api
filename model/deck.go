package model

import (
	"github.com/raulcbs/flipcard-api/database"
	"gorm.io/gorm"
)

type Deck struct {
	gorm.Model
	Name        string `gorm:"size:255;not null" json:"name"`
	NmUserDecks []NmUserDeck
	Cards       []Card
}

func (deck *Deck) Save() (*Deck, error) {
	err := database.Database.Create(&deck).Error
	if err != nil {
		return &Deck{}, err
	}
	return deck, nil
}

func (deck *Deck) GetAll() ([]Deck, error) {
	var decks []Deck
	err := database.Database.Find(&decks).Error
	if err != nil {
		return nil, err
	}
	return decks, nil
}
