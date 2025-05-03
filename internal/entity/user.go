package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64
	Username  string
	FullName  string
	Email     string
	Password  string
	Role      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
