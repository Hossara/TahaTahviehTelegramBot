package port

import (
	"io"
	"taha_tahvieh_tg_bot/internal/product_storage/domain"

	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
	storageDomain "taha_tahvieh_tg_bot/internal/storage/domain"
)

type ClientRepo interface {
	BucketExists(name string) (bool, error)
	CreateBucket(name string) error

	UploadFile(file *domain.File, url string) error
	StreamFile(bucket, name string) (io.ReadCloser, error)

	DeleteFile(bucket string, filePath string) error
	DeleteMultipleFiles(bucket string, filePaths []string) error
}

type Repo interface {
	Insert(file *domain.File) (storageDomain.FileID, error)

	FindAllFilesByProductID(productID productDomain.ProductID) ([]*domain.File, error)

	DeleteFileByID(fileID storageDomain.FileID) error

	DeleteAllFilesByProductID(productID productDomain.ProductID) error

	RunMigrations() error
}
