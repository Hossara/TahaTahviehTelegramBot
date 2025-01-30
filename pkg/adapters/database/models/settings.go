package models

import (
	"time"
)

type Setting struct {
	ID        uint8     `gorm:"primarykey;autoIncrement"`
	Title     string    `gorm:"size:255"`
	Content   string    `gorm:"size:255;type:json;default:'{}'"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
