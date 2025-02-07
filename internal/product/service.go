package product

import (
	"context"
	"errors"
	"fmt"
	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
	"taha_tahvieh_tg_bot/internal/product/port"
	"taha_tahvieh_tg_bot/internal/product_storage/domain"
	storagePort "taha_tahvieh_tg_bot/internal/storage/port"
)

var (
	ErrPaginationInvalid  = errors.New("invalid pagination parameters")
	ErrBrandOrTypeInvalid = errors.New("brand or type invalid")
	ErrInvalidID          = errors.New("invalid id")
)

type service struct {
	repo     port.Repo
	metaRepo port.MetaRepo
	storage  storagePort.Service
	ctx      context.Context
}

func NewService(ctx context.Context, repo port.Repo, metaRepo port.MetaRepo, storage storagePort.Service) port.Service {
	return &service{
		ctx:      ctx,
		repo:     repo,
		metaRepo: metaRepo,
		storage:  storage,
	}
}

func (r *service) CreateProduct(product *domain.Product) (productDomain.ProductID, error) {
	if product.Brand.ID == 0 || product.Type.ID == 0 {
		return 0, ErrBrandOrTypeInvalid
	}

	id, err := r.repo.Insert(product)

	if err != nil {
		return 0, fmt.Errorf("error while insert product: %w", err)
	}

	return id, nil
}

func (r *service) GetAllProductsBasedOn(
	brandID productDomain.BrandID,
	productTypeID productDomain.ProductTypeID,
	title string, page, pageSize int,
) (domain.ProductPagination, error) {
	if page < 1 || pageSize < 1 {
		return domain.ProductPagination{}, ErrPaginationInvalid
	}

	var result domain.ProductPagination
	var err error

	if title != "" {
		result, err = r.repo.FindAllByTitle(title, true, page, pageSize)

		if err != nil {
			return domain.ProductPagination{}, fmt.Errorf("error fetching products by title: %w", err)
		}

		return result, nil
	}

	// If no title, search by brand/type
	result, err = r.repo.FindAllByMeta(brandID, productTypeID, true, page, pageSize)

	if err != nil {
		return domain.ProductPagination{}, fmt.Errorf("error fetching products by meta filters: %w", err)
	}

	return result, nil
}

func (r *service) GetAllBrands(page, pageSize int) (productDomain.BrandPagination, error) {
	if page < 1 || pageSize < 1 {
		return productDomain.BrandPagination{}, ErrPaginationInvalid
	}

	result, err := r.metaRepo.FindAllBrand(page, pageSize)
	if err != nil {
		return productDomain.BrandPagination{}, fmt.Errorf("failed to fetch brands: %w", err)
	}

	return result, nil
}

func (r *service) GetAllProductTypes(page, pageSize int) (productDomain.ProductTypePagination, error) {
	if page < 1 || pageSize < 1 {
		return productDomain.ProductTypePagination{}, ErrPaginationInvalid
	}

	result, err := r.metaRepo.FindAllProductType(page, pageSize)

	if err != nil {
		return productDomain.ProductTypePagination{}, fmt.Errorf("failed to fetch product types: %w", err)
	}

	return result, nil
}

func (r *service) GetProduct(id productDomain.ProductID) (*domain.Product, error) {
	if id == 0 {
		return nil, ErrInvalidID
	}

	product, err := r.repo.FindByID(id, true)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch product: %w", err)
	}

	return product, nil
}

func (r *service) CreateBrand(brand *productDomain.Brand) error {
	if brand == nil {
		return errors.New("brand cannot be empty")
	}

	err := r.metaRepo.InsertBrand(brand)

	if err != nil {
		return fmt.Errorf("error inserting brand: %w", err)
	}

	return nil
}

func (r *service) CreateProductType(productType *productDomain.ProductType) error {
	if productType == nil {
		return errors.New("product type cannot be empty")
	}

	err := r.metaRepo.InsertProductType(productType)

	if err != nil {
		return fmt.Errorf("error inserting product type: %w", err)
	}

	return nil
}

func (r *service) DeleteProductType(id productDomain.ProductTypeID) error {
	if id == 0 {
		return ErrInvalidID
	}

	err := r.metaRepo.DeleteProductType(id)

	if err != nil {
		return fmt.Errorf("error deleting product type: %w", err)
	}

	return nil
}

func (r *service) DeleteBrand(id productDomain.BrandID) error {
	if id == 0 {
		return ErrInvalidID
	}

	err := r.metaRepo.DeleteBrand(id)

	if err != nil {
		return fmt.Errorf("error deleting brand: %w", err)
	}

	return nil
}

func (r *service) DeleteProduct(id productDomain.ProductID, files []domain.File) error {
	if id == 0 {
		return ErrInvalidID
	}

	err := r.storage.RemoveAllFiles(files)

	if err != nil {
		return fmt.Errorf("error deleting product files: %w", err)
	}

	err = r.repo.DeleteById(id)

	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}

	return nil
}

func (r *service) UpdateProduct(id productDomain.ProductID, updates map[string]interface{}) error {
	if id == 0 {
		return ErrInvalidID
	}

	if len(updates) == 0 {
		return nil
	}

	err := r.repo.UpdateByID(id, updates)

	if err != nil {
		return fmt.Errorf("error updating product: %w", err)
	}
	return nil
}

func (r *service) UpdateBrand(id productDomain.BrandID, updates map[string]interface{}) error {
	if id == 0 {
		return ErrInvalidID
	}

	if len(updates) == 0 {
		return nil
	}

	err := r.metaRepo.UpdateBrand(id, updates)

	if err != nil {
		return fmt.Errorf("error updating brand: %w", err)
	}
	return nil
}

func (r *service) UpdateProductType(id productDomain.ProductTypeID, updates map[string]interface{}) error {
	if id == 0 {
		return ErrInvalidID
	}

	if len(updates) == 0 {
		return nil
	}

	err := r.metaRepo.UpdateProductType(id, updates)

	if err != nil {
		return fmt.Errorf("error updating product type: %w", err)
	}
	return nil
}

func (r *service) RunProductMigrations() error {
	return r.repo.RunMigrations()
}

func (r *service) RunProductTypeMigrations() error {
	return r.metaRepo.RunProductTypeMigrations()
}

func (r *service) RunBrandMigrations() error {
	return r.metaRepo.RunBrandMigrations()
}
