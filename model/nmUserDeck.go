package model

import (
	"github.com/raulcbs/flipcard-api/database"
	"gorm.io/gorm"
)

type NmUserDeck struct {
	gorm.Model
	UserID uint
	DeckID uint
}

func (nm *NmUserDeck) Save() (*NmUserDeck, error) {
	err := database.Database.Create(&nm).Error
	if err != nil {
		return &NmUserDeck{}, err
	}
	return nm, nil
}
