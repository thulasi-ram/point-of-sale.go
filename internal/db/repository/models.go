// Code generated by sqlc. DO NOT EDIT.

package repository

import (
	"database/sql"
	"time"
)

type Product struct {
	ID          int64
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CategoryID  int64
}

type ProductCategory struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	ParentID  sql.NullInt64
}

type Supplier struct {
	ID        int64
	Name      string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
