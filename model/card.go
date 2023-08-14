package model

import (
	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	Front  string `gorm:"not null"`
	Back   string `gorm:"not null"`
	DeckID uint
}
