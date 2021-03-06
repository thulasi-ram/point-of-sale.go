// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: products.queries.sql

package repository

import (
	"context"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (name, description, category_id)
VALUES ($1, $2, $3)
RETURNING id, name, description, created_at, updated_at, category_id
`

type CreateProductParams struct {
	Name        string
	Description string
	CategoryID  int64
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct, arg.Name, arg.Description, arg.CategoryID)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CategoryID,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE
FROM products
WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteProduct, id)
	return err
}

const getProduct = `-- name: GetProduct :one
SELECT id, name, description, created_at, updated_at, category_id
FROM products
WHERE id = $1
`

func (q *Queries) GetProduct(ctx context.Context, id int64) (Product, error) {
	row := q.db.QueryRow(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CategoryID,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT id, name, description, created_at, updated_at, category_id
FROM products
ORDER BY name
`

func (q *Queries) ListProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, listProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CategoryID,
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
