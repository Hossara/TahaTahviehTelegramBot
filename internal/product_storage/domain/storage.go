package domain

import (
	"github.com/google/uuid"
	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
	storageDomain "taha_tahvieh_tg_bot/internal/storage/domain"
)

type File struct {
	ID          storageDomain.FileID
	ProductID   productDomain.ProductID
	UUID        uuid.UUID
	BucketName  string
	Path        string
	Format      string
	Size        int64
	ContentType storageDomain.ContentType
}
