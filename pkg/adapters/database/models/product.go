package models

import "github.com/google/uuid"

type Product struct {
	ID          int64       `gorm:"primaryKey;autoIncrement"`
	UUID        uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title       string      `gorm:"not null"`
	BrandID     int64       `gorm:"not null"` // Foreign key to Brand
	Brand       Brand       `gorm:"foreignKey:BrandID"`
	TypeID      int64       `gorm:"not null"` // Foreign key to ProductType
	Type        ProductType `gorm:"foreignKey:TypeID"`
	Description string      `gorm:"type:text"`

	Files []File `gorm:"foreignKey:ProductID"` // Relationship to files
}
