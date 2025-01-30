package port

import (
	"taha_tahvieh_tg_bot/internal/product_storage/domain"

	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
)

type Service interface {
	InitBucket(name string) error

	UploadFile(file *domain.File) error

	RemoveFile(filePath string) error
	RemoveAllProductFiles(productID productDomain.ProductID) error

	GetProductFiles(productID productDomain.ProductID) ([]*domain.File, error)

	RunMigrations() error
}
