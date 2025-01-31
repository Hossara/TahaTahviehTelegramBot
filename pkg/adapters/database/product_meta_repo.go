package database

import (
	"errors"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
	"taha_tahvieh_tg_bot/internal/product/port"
	"taha_tahvieh_tg_bot/pkg/adapters/database/helpers"
	"taha_tahvieh_tg_bot/pkg/adapters/database/mapper"
	"taha_tahvieh_tg_bot/pkg/adapters/database/models"
	"taha_tahvieh_tg_bot/pkg/utils"
)

var (
	ErrBrandNotFound       = errors.New("brand no found")
	ErrProductTypeNotFound = errors.New("product type no found")
)

type productMetaRepo struct {
	db *gorm.DB
}

func NewProductMetaRepo(db *gorm.DB) port.MetaRepo {
	return &productMetaRepo{db}
}

func (r *productMetaRepo) FindAllBrand(page int, pageSize int) (productDomain.BrandPagination, error) {
	var brands []*models.Brand
	var total int64

	r.db.Model(&models.Brand{}).Count(&total)
	offset := (page - 1) * pageSize

	return productDomain.BrandPagination{
		Page:  page,
		Pages: int((total + int64(pageSize) - 1) / int64(pageSize)),
		Data: utils.Map(brands, func(t *models.Brand) *productDomain.Brand {
			return mapper.ToDomainBrand(t)
		}),
	}, r.db.Limit(pageSize).Offset(offset).Find(&brands).Error
}

// FindAllProductType retrieves all product types
func (r *productMetaRepo) FindAllProductType(page int, pageSize int) (productDomain.ProductTypePagination, error) {
	var pts []*models.ProductType
	var total int64

	r.db.Model(&models.ProductType{}).Count(&total)
	offset := (page - 1) * pageSize

	return productDomain.ProductTypePagination{
		Page:  page,
		Pages: int((total + int64(pageSize) - 1) / int64(pageSize)),
		Data: utils.Map(pts, func(t *models.ProductType) *productDomain.ProductType {
			return mapper.ToDomainProductType(t)
		}),
	}, r.db.Limit(pageSize).Offset(offset).Find(&pts).Error
}

// InsertBrand inserts a new brand
func (r *productMetaRepo) InsertBrand(brand *productDomain.Brand) error {
	modelBrand := mapper.ToModelBrand(brand)

	return r.db.Create(modelBrand).Error
}

// UpdateBrand updates a brand partially by ID
func (r *productMetaRepo) UpdateBrand(id productDomain.BrandID, updates map[string]interface{}) error {
	result := r.db.Model(&models.Brand{}).Where("id = ?", id).Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrBrandNotFound
	}
	return nil
}

// DeleteBrand deletes a brand by ID
func (r *productMetaRepo) DeleteBrand(id productDomain.BrandID) error {
	result := r.db.Delete(&models.Brand{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrBrandNotFound
	}
	return nil
}

// InsertProductType inserts a new product type
func (r *productMetaRepo) InsertProductType(productType *productDomain.ProductType) error {
	modelProductType := mapper.ToModelProductType(productType)

	return r.db.Create(modelProductType).Error
}

// UpdateProductType updates a product type partially by ID
func (r *productMetaRepo) UpdateProductType(id productDomain.ProductTypeID, updates map[string]interface{}) error {
	result := r.db.Model(&models.ProductType{}).Where("id = ?", id).Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrProductTypeNotFound
	}
	return nil
}

// DeleteProductType deletes a product type by ID
func (r *productMetaRepo) DeleteProductType(id productDomain.ProductTypeID) error {
	result := r.db.Delete(&models.ProductType{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrProductTypeNotFound
	}
	return nil
}

func (r *productMetaRepo) RunProductTypeMigrations() error {
	migrator := gormigrate.New(
		r.db, gormigrate.DefaultOptions,
		helpers.GetMigrations[models.ProductType]("product_types", &models.ProductType{}),
	)

	return migrator.Migrate()
}

func (r *productMetaRepo) RunBrandMigrations() error {
	migrator := gormigrate.New(
		r.db, gormigrate.DefaultOptions,
		helpers.GetMigrations[models.Product]("brands", &models.Product{}),
	)

	return migrator.Migrate()
}
