package port

import (
	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
	"taha_tahvieh_tg_bot/internal/product_storage/domain"
)

type Repo interface {
	FindAll(preload bool, page, pageSize int) (domain.ProductPagination, error)
	FindAllByTitle(title string, preload bool, page, pageSize int) (domain.ProductPagination, error)
	FindByID(id productDomain.ProductID, preload bool) (*domain.Product, error)
	FindAllByMeta(
		brandID productDomain.BrandID, productTypeID productDomain.ProductTypeID,
		preload bool, page, pageSize int,
	) (domain.ProductPagination, error)

	UpdateByID(id productDomain.ProductID, updates map[string]interface{}) error

	Insert(product *domain.Product) (productDomain.ProductID, error)

	DeleteById(id productDomain.ProductID) error

	RunMigrations() error
}

type MetaRepo interface {
	FindAllBrand(page, pageSize int) (productDomain.BrandPagination, error)
	FindAllProductType(page, pageSize int) (productDomain.ProductTypePagination, error)

	InsertBrand(brand *productDomain.Brand) error
	UpdateBrand(id productDomain.BrandID, updates map[string]interface{}) error
	DeleteBrand(id productDomain.BrandID) error

	InsertProductType(brand *productDomain.ProductType) error
	UpdateProductType(id productDomain.ProductTypeID, updates map[string]interface{}) error
	DeleteProductType(id productDomain.ProductTypeID) error

	RunProductTypeMigrations() error
	RunBrandMigrations() error
}
