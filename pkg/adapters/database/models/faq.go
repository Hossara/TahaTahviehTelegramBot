package models

import "time"

type Faq struct {
	ID        uint8     `gorm:"primarykey;autoIncrement"`
	Question  string    `gorm:"size:255"`
	Answer    string    `gorm:"size:255"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
