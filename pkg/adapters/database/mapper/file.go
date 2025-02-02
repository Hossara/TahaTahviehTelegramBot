package mapper

import (
	"taha_tahvieh_tg_bot/internal/product_storage/domain"
	"taha_tahvieh_tg_bot/pkg/adapters/database/models"

	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
	storageDomain "taha_tahvieh_tg_bot/internal/storage/domain"
)

func ToDomainFile(m *models.File) *domain.File {
	return &domain.File{
		ID:          storageDomain.FileID(m.ID),
		BucketName:  m.BucketName,
		UUID:        m.UUID,
		Path:        m.Path,
		Format:      m.Format,
		Size:        m.Size,
		ContentType: storageDomain.ContentType(m.ContentType),
		ProductID:   productDomain.ProductID(m.ProductID),
	}
}

func ToModelFile(d *domain.File) *models.File {
	return &models.File{
		ID:          int64(d.ID),
		BucketName:  d.BucketName,
		UUID:        d.UUID,
		Path:        d.Path,
		Format:      d.Format,
		Size:        d.Size,
		ContentType: string(d.ContentType),
		ProductID:   int64(d.ProductID),
	}
}
