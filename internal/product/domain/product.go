package domain

import (
	"taha_tahvieh_tg_bot/internal/common"
)

type BrandID int64
type ProductTypeID int64
type ProductID int64

type Brand struct {
	ID          BrandID
	Title       string
	Description string
}

type ProductType struct {
	ID          ProductTypeID
	Title       string
	Description string
}

type ProductTypePagination common.Pagination[[]*ProductType]
type BrandPagination common.Pagination[[]*Brand]
