package entity

import (
	"database/sql"
	"time"
)

type Supplier struct {
	ID            int64
	Name          string
	ContactPerson *string
	Phone         *string
	Address       *string
	Email         *string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime
}
