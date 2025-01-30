package mapper

import (
	"taha_tahvieh_tg_bot/internal/product_storage/domain"
	"taha_tahvieh_tg_bot/pkg/adapters/database/models"
	"taha_tahvieh_tg_bot/pkg/utils"

	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
)

func ToDomainProduct(m *models.Product) *domain.Product {
	return &domain.Product{
		ID:          productDomain.ProductID(m.ID),
		UUID:        m.UUID,
		Title:       m.Title,
		Description: m.Description,
		Brand:       *ToDomainBrand(&m.Brand),
		Type:        *ToDomainProductType(&m.Type),
		Files: utils.Map(m.Files, func(t models.File) domain.File {
			return *ToDomainFile(&t)
		}),
	}
}

func ToModelProduct(d *domain.Product) *models.Product {
	return &models.Product{
		ID:          int64(d.ID),
		UUID:        d.UUID,
		Title:       d.Title,
		Description: d.Description,
		BrandID:     int64(d.Brand.ID),
		Brand:       *ToModelBrand(&d.Brand),
		TypeID:      int64(d.Type.ID),
		Type:        *ToModelProductType(&d.Type),
		Files: utils.Map(d.Files, func(t domain.File) models.File {
			return *ToModelFile(&t)
		}),
	}
}
