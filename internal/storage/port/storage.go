package port

import (
	"io"
	"taha_tahvieh_tg_bot/internal/product_storage/domain"
	storageDomain "taha_tahvieh_tg_bot/internal/storage/domain"

	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
)

type Service interface {
	InitBucket(name string) error

	UploadFile(file *domain.File, url string) (storageDomain.FileID, error)

	RemoveFile(filePath string) error
	RemoveAllFiles(files []domain.File) error
	RemoveAllProductFiles(productID productDomain.ProductID, files []domain.File) error

	GetProductFile(files domain.File) (io.ReadCloser, error)

	RunMigrations() error
}
