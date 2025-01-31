package domain

import (
	"github.com/google/uuid"
	"taha_tahvieh_tg_bot/internal/common"
	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
)

type Product struct {
	ID          productDomain.ProductID
	UUID        uuid.UUID
	Title       string
	Brand       productDomain.Brand
	Type        productDomain.ProductType
	Description string
	Files       []File
}

type ProductPagination common.Pagination[[]*Product]
