package model

import (
	"time"

	"gorm.io/gorm"
)

type Word struct {
	gorm.Model
	ID              uint      `gorm:"primaryKey" json:"id"`
	Value           string    `gorm:"size:255;not null" json:"value"`
	ExampleSentence string    `gorm:"column:example_sentence;type:text" json:"example_sentence"`
	Meaning         string    `gorm:"size:255;not null" json:"meaning"`
	Audio           string    `json:"audio"`
	Image           string    `json:"image"`
	CreatedAt       time.Time `gorm:"not null" json:"-"`
	UpdatedAt       time.Time `gorm:"not null" json:"-"`
	CardID          uint
}
