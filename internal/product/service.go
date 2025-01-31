package product

import (
	"context"
	"errors"
	"fmt"
	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
	"taha_tahvieh_tg_bot/internal/product/port"
	"taha_tahvieh_tg_bot/internal/product_storage/domain"
)

var (
	ErrPaginationInvalid = errors.New("invalid pagination parameters")
)

type service struct {
	repo     port.Repo
	metaRepo port.MetaRepo
	ctx      context.Context
}

func NewService(ctx context.Context, repo port.Repo, metaRepo port.MetaRepo) port.Service {
	return &service{
		ctx:      ctx,
		repo:     repo,
		metaRepo: metaRepo,
	}
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
	//TODO implement me
	panic("implement me")
}

func (r *service) CreateBrand(brand *productDomain.Brand) error {
	//TODO implement me
	panic("implement me")
}

func (r *service) CreateProductType(productType productDomain.ProductType) error {
	//TODO implement me
	panic("implement me")
}

func (r *service) DeleteProductType(id productDomain.ProductTypeID) error {
	//TODO implement me
	panic("implement me")
}

func (r *service) DeleteBrand(id productDomain.BrandID) error {
	//TODO implement me
	panic("implement me")
}

func (r *service) DeleteProduct(id productDomain.ProductID) error {
	//TODO implement me
	panic("implement me")
}

func (r *service) UpdateProduct(id productDomain.ProductID, updates map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (r *service) UpdateBrand(id productDomain.BrandID, updates map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (r *service) UpdateProductType(id productDomain.ProductTypeID, updates map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
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
