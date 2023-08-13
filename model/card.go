package model

import (
	"time"

	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	ID        uint   `orm:"primaryKey;not null" json:"id"`
	Front     string `gorm:"not null" json:"front"`
	Back      string `gorm:"not null" json:"back"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeckID    uint
	Words     []Word
}
