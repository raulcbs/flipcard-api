package model

import (
	"time"

	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	ID        uint   `orm:"primaryKey;not null" json:"id"`
	TypeCard  string `orm:"column:type_card;size:255;not null" json:"type_card"`
	Mode      string `orm:"size:255" json:"mode"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeckID    uint
	Word      Word
}
