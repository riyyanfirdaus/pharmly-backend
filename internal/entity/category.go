package entity

import (
	"database/sql"
	"time"
)

type Category struct {
	ID               int64
	Name             string
	Description      string
	ParentCategoryID *int64
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        sql.NullTime
}
