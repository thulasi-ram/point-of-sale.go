// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: customer.queries.sql

package repository

import (
	"context"
)

const createCustomer = `-- name: CreateCustomer :one
INSERT INTO customers (name, phone, address)
VALUES ($1, $2, $3)
RETURNING id, name, phone, address, created_at, updated_at
`

type CreateCustomerParams struct {
	Name    string
	Phone   string
	Address string
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error) {
	row := q.db.QueryRow(ctx, createCustomer, arg.Name, arg.Phone, arg.Address)
	var i Customer
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

const deleteCustomer = `-- name: DeleteCustomer :exec
DELETE
FROM customers
WHERE id = $1
`

func (q *Queries) DeleteCustomer(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteCustomer, id)
	return err
}

const getCustomer = `-- name: GetCustomer :one
SELECT id, name, phone, address, created_at, updated_at
FROM customers
WHERE id = $1
`

func (q *Queries) GetCustomer(ctx context.Context, id int64) (Customer, error) {
	row := q.db.QueryRow(ctx, getCustomer, id)
	var i Customer
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

const listCustomers = `-- name: ListCustomers :many
SELECT id, name, phone, address, created_at, updated_at
FROM customers
ORDER BY name
`

func (q *Queries) ListCustomers(ctx context.Context) ([]Customer, error) {
	rows, err := q.db.Query(ctx, listCustomers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Customer
	for rows.Next() {
		var i Customer
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
