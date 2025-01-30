package models

import "github.com/google/uuid"

type File struct {
	ID          int64     `gorm:"primaryKey;autoIncrement"`   // Primary key with auto-increment
	BucketName  string    `gorm:"type:varchar(255);not null"` // Alternate unique identifier
	UUID        uuid.UUID `gorm:"type:uuid;not null"`         // Unique identifier, type UUID
	Path        string    `gorm:"type:text;not null"`         // File storage path
	Format      string    `gorm:"type:varchar(50);not null"`  // File format (e.g., png, pdf)
	Size        int64     `gorm:"not null"`                   // File size in bytes
	ContentType string    `gorm:"type:varchar(50);not null"`  // MIME type for content

	ProductID int64   `gorm:"not null"`
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ProductID"` // Relation to Product
}
