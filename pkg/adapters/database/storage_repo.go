package database

import (
	"errors"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"taha_tahvieh_tg_bot/internal/product_storage/domain"
	storageDomain "taha_tahvieh_tg_bot/internal/storage/domain"
	"taha_tahvieh_tg_bot/internal/storage/port"
	"taha_tahvieh_tg_bot/pkg/adapters/database/helpers"
	"taha_tahvieh_tg_bot/pkg/adapters/database/mapper"
	"taha_tahvieh_tg_bot/pkg/adapters/database/models"
	"taha_tahvieh_tg_bot/pkg/utils"

	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
)

var (
	ErrFileNotFound         = errors.New("file no found")
	ErrProductFilesNotFound = errors.New("product file no found")
)

type storageRepo struct {
	db *gorm.DB
}

func NewStorageRepo(db *gorm.DB) port.Repo {
	return &storageRepo{db}
}

func (r *storageRepo) Insert(file *domain.File) (storageDomain.FileID, error) {
	modelFile := mapper.ToModelFile(file)

	return storageDomain.FileID(modelFile.ID), r.db.Create(&modelFile).Error
}

func (r *storageRepo) FindAllFilesByProductID(productID productDomain.ProductID) ([]*domain.File, error) {
	var files []models.File

	return utils.Map(files, func(t models.File) *domain.File {
		return mapper.ToDomainFile(&t)
	}), r.db.Where("product_id = ?", int64(productID)).Find(&files).Error
}

func (r *storageRepo) DeleteFileByID(fileID storageDomain.FileID) error {
	result := r.db.Delete(&models.File{}, int64(fileID))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrFileNotFound
	}

	return nil
}

func (r *storageRepo) DeleteAllFilesByProductID(productID productDomain.ProductID) error {
	result := r.db.Where("product_id = ?", int64(productID)).Delete(&models.File{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrProductFilesNotFound
	}
	return nil
}

func (r *storageRepo) RunMigrations() error {
	migrator := gormigrate.New(
		r.db, gormigrate.DefaultOptions,
		helpers.GetMigrations[models.File]("files", &models.File{}),
	)

	return migrator.Migrate()
}
