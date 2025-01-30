package database

import (
	"errors"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"taha_tahvieh_tg_bot/internal/common"
	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
	"taha_tahvieh_tg_bot/internal/product/port"
	"taha_tahvieh_tg_bot/internal/product_storage/domain"
	"taha_tahvieh_tg_bot/pkg/adapters/database/helpers"
	"taha_tahvieh_tg_bot/pkg/adapters/database/mapper"
	"taha_tahvieh_tg_bot/pkg/adapters/database/models"
	"taha_tahvieh_tg_bot/pkg/utils"
)

var (
	ErrProductNotFound = errors.New("product no found")
)

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) port.Repo {
	return &productRepo{db}
}

func (r *productRepo) FindAllByMeta(
	brandID productDomain.BrandID, productTypeID productDomain.ProductTypeID,
	preload bool, page int, pageSize int,
) (common.Pagination, error) {
	var products []*models.Product
	var total int64

	query := r.db.Where("brand_id = ? AND type_id = ?", brandID, productTypeID)

	if preload {
		query = query.Preload("Brand").Preload("Type").Preload("Files")
	}

	// Get total count for pagination
	query.Model(&models.Product{}).Count(&total)

	// Apply pagination
	offset := (page - 1) * pageSize

	return common.Pagination{
		Page:  page,
		Pages: int((total + int64(pageSize) - 1) / int64(pageSize)),
		Data: utils.Map(products, func(t *models.Product) *domain.Product {
			return mapper.ToDomainProduct(t)
		}),
	}, query.Limit(pageSize).Offset(offset).Find(&products).Error
}

func (r *productRepo) FindAll(preload bool, page int, pageSize int) (common.Pagination, error) {
	var products []*models.Product
	var total int64
	db := r.db

	if preload {
		db = db.Preload("Brand").Preload("Type").Preload("Files")
	}

	db.Model(&models.Product{}).Count(&total)
	offset := (page - 1) * pageSize

	return common.Pagination{
		Page:  page,
		Pages: int((total + int64(pageSize) - 1) / int64(pageSize)),
		Data: utils.Map(products, func(t *models.Product) *domain.Product {
			return mapper.ToDomainProduct(t)
		}),
	}, db.Limit(pageSize).Offset(offset).Find(&products).Error
}

// FindByID retrieves a product by ID
func (r *productRepo) FindByID(id productDomain.ProductID, preload bool) (*domain.Product, error) {
	var product models.Product
	db := r.db

	if preload {
		db = db.Preload("Brand").Preload("Type").Preload("Files")
	}

	return mapper.ToDomainProduct(&product), db.First(&product, id).Error
}

// UpdateByID updates a product partially by ID
func (r *productRepo) UpdateByID(id productDomain.ProductID, updates map[string]interface{}) error {
	result := r.db.Model(&models.Product{}).Where("id = ?", id).Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrProductNotFound
	}
	return nil
}

// Insert inserts a new product
func (r *productRepo) Insert(product *domain.Product) error {
	modelProduct := mapper.ToModelProduct(product)

	return r.db.Create(modelProduct).Error
}

// DeleteById deletes a product by ID
func (r *productRepo) DeleteById(id productDomain.ProductID) error {
	result := r.db.Delete(&models.Product{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrProductNotFound
	}
	return nil
}

func (r *productRepo) RunMigrations() error {
	migrator := gormigrate.New(
		r.db, gormigrate.DefaultOptions,
		helpers.GetMigrations[models.Product]("products", &models.Product{}),
	)

	return migrator.Migrate()
}
