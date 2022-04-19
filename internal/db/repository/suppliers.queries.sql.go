// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: suppliers.queries.sql

package repository

import (
	"context"
)

const createSupplier = `-- name: CreateSupplier :one
INSERT INTO suppliers (name, phone, address)
VALUES ($1, $2, $3)
RETURNING id, name, phone, address, created_at, updated_at
`

type CreateSupplierParams struct {
	Name    string
	Phone   string
	Address string
}

func (q *Queries) CreateSupplier(ctx context.Context, arg CreateSupplierParams) (Supplier, error) {
	row := q.db.QueryRow(ctx, createSupplier, arg.Name, arg.Phone, arg.Address)
	var i Supplier
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Phone,
		&i.Address,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteSupplier = `-- name: DeleteSupplier :exec
DELETE
FROM suppliers
WHERE id = $1
`

func (q *Queries) DeleteSupplier(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteSupplier, id)
	return err
}

const getSupplier = `-- name: GetSupplier :one
SELECT id, name, phone, address, created_at, updated_at
FROM suppliers
WHERE id = $1
`

func (q *Queries) GetSupplier(ctx context.Context, id int64) (Supplier, error) {
	row := q.db.QueryRow(ctx, getSupplier, id)
	var i Supplier
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Phone,
		&i.Address,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listSuppliers = `-- name: ListSuppliers :many
SELECT id, name, phone, address, created_at, updated_at
FROM suppliers
ORDER BY name
`

func (q *Queries) ListSuppliers(ctx context.Context) ([]Supplier, error) {
	rows, err := q.db.Query(ctx, listSuppliers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Supplier
	for rows.Next() {
		var i Supplier
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Phone,
			&i.Address,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
