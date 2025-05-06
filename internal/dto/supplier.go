package dto

import (
	"database/sql"
	"time"
)

type SupplierRequest struct {
	Name          string `json:"name,omitempty"`
	ContactPerson string `json:"contact_person,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Address       string `json:"address,omitempty"`
	Email         string `json:"email,omitempty"`
}

type SupplierResponse struct {
	ID            int64        `json:"id"`
	Name          string       `json:"name"`
	ContactPerson string       `json:"contact_person"`
	Phone         string       `json:"phone"`
	Address       string       `json:"address"`
	Email         string       `json:"email"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	DeletedAt     sql.NullTime `json:"deleted_at"`
}
