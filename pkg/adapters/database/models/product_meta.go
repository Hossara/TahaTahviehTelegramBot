package models

type ProductType struct {
	ID          int64  `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"not null"`
	Description string `gorm:"type:text"`
}

type Brand struct {
	ID          int64  `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"not null"`
	Description string `gorm:"type:text"`
}
