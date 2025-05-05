package dto

import (
	"database/sql"
	"time"
)

type CategoryRequest struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	ParentCategoryID int64  `json:"parent_category_id,omitempty"`
}

type CategoryResponse struct {
	ID               int64        `json:"id"`
	Name             string       `json:"name"`
	Description      string       `json:"description"`
	ParentCategoryID int64        `json:"parent_category_id"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at"`
	DeletedAt        sql.NullTime `json:"deleted_at"`
}
