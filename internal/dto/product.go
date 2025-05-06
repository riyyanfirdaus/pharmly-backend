package dto

import (
	"database/sql"
	"time"

	"github.com/shopspring/decimal"
)

type ProductRequest struct {
	Name           string          `json:"name"`
	CategoryID     int64           `json:"category_id"`
	GenericName    string          `json:"generic_name"`
	Description    string          `json:"description,omitempty"`
	Price          decimal.Decimal `json:"price"`
	Stock          int             `json:"stock"`
	Unit           string          `json:"unit"`
	ExpirationDate time.Time       `json:"expiration_date"`
	Barcode        string          `json:"barcode"`
	SupplierID     int64           `json:"supplier_id"`
	MinStock       int             `json:"min_stock"`
	IsActive       bool            `json:"is_active,omitempty"`
}

type ProductResponse struct {
	ID             int64           `json:"id"`
	Name           string          `json:"name"`
	CategoryID     int64           `json:"category_id"`
	GenericName    string          `json:"generic_name"`
	Description    *string         `json:"description"`
	Price          decimal.Decimal `json:"price"`
	Stock          int             `json:"stock"`
	Unit           string          `json:"unit"`
	ExpirationDate time.Time       `json:"expiration_date"`
	Barcode        string          `json:"barcode"`
	SupplierID     int64           `json:"supplier_id"`
	MinStock       int             `json:"min_stock"`
	IsActive       bool            `json:"is_active"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeletedAt      sql.NullTime    `json:"deleted_at"`
}
