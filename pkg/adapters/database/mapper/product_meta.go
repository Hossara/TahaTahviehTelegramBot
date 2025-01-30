package mapper

import (
	"taha_tahvieh_tg_bot/internal/product/domain"
	"taha_tahvieh_tg_bot/pkg/adapters/database/models"
)

func ToModelBrand(d *domain.Brand) *models.Brand {
	return &models.Brand{
		ID:          int64(d.ID),
		Title:       d.Title,
		Description: d.Description,
	}
}

func ToDomainBrand(m *models.Brand) *domain.Brand {
	return &domain.Brand{
		ID:          domain.BrandID(m.ID),
		Title:       m.Title,
		Description: m.Description,
	}
}

func ToModelProductType(d *domain.ProductType) *models.ProductType {
	return &models.ProductType{
		ID:          int64(d.ID),
		Title:       d.Title,
		Description: d.Description,
	}
}

func ToDomainProductType(m *models.ProductType) *domain.ProductType {
	return &domain.ProductType{
		ID:          domain.ProductTypeID(m.ID),
		Title:       m.Title,
		Description: m.Description,
	}
}
