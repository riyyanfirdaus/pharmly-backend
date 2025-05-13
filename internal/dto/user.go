package dto

import (
	"database/sql"
	"time"
)

type UserRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	FullName string `json:"full_name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"  validate:"required,min=6"`
	Role     string `json:"role"  validate:"required,oneof=admin pharmacist cashier"`
}

type UserResponse struct {
	ID        int64        `json:"id"`
	Username  string       `json:"username"`
	FullName  string       `json:"full_name"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	Role      string       `json:"role"`
	Status    string       `json:"status"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Token string        `json:"token"`
	User  *UserResponse `json:"user"`
}
