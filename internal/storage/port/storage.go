package port

import (
	"taha_tahvieh_tg_bot/internal/product_storage/domain"
	storageDomain "taha_tahvieh_tg_bot/internal/storage/domain"

	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
)

type Service interface {
	InitBucket(name string) error

	UploadFile(file *domain.File, url string) (storageDomain.FileID, error)

	RemoveFile(filePath string) error
	RemoveAllFiles(files []domain.File) error

	GetProductFiles(productID productDomain.ProductID) ([]*domain.File, error)

	RunMigrations() error
}
