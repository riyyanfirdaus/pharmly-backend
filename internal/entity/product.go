package entity

import (
	"database/sql"
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID             int64
	Name           string
	CategoryID     int64
	GenericName    string
	Description    *string
	Price          decimal.Decimal
	Stock          int
	Unit           string
	ExpirationDate time.Time
	Barcode        string
	SupplierID     int64
	MinStock       int
	IsActive       bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime
}
