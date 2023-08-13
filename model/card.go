package model

import (
	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	Front  string `gorm:"not null" json:"front"`
	Back   string `gorm:"not null" json:"back"`
	DeckID uint
	Words  []Word
}
