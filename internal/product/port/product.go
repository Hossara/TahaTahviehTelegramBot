package port

import (
	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
	"taha_tahvieh_tg_bot/internal/product_storage/domain"
)

type Service interface {
	GetAllProductsBasedOn(
		brandID productDomain.BrandID,
		productTypeID productDomain.ProductTypeID,
		title string,
		page, pageSize int,
	) (domain.ProductPagination, error)

	GetAllBrands(page, pageSize int) (productDomain.BrandPagination, error)
	GetAllProductTypes(page, pageSize int) (productDomain.ProductTypePagination, error)

	GetProduct(id productDomain.ProductID) (*domain.Product, error)

	CreateProduct(product *domain.Product) (productDomain.ProductID, error)
	CreateBrand(brand *productDomain.Brand) error
	CreateProductType(productType *productDomain.ProductType) error

	DeleteProductType(id productDomain.ProductTypeID) error
	DeleteBrand(id productDomain.BrandID) error
	DeleteProduct(id productDomain.ProductID) error

	UpdateProduct(id productDomain.ProductID, updates map[string]interface{}) error
	UpdateBrand(id productDomain.BrandID, updates map[string]interface{}) error
	UpdateProductType(id productDomain.ProductTypeID, updates map[string]interface{}) error

	RunProductMigrations() error
	RunProductTypeMigrations() error
	RunBrandMigrations() error
}
