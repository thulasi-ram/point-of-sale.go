// Code generated by sqlc. DO NOT EDIT.
// source: product_categories.queries.sql

package repository

import (
	"context"
	"database/sql"
)

const createProductCategory = `-- name: CreateProductCategory :one
INSERT INTO product_categories (name, parent_id)
VALUES ($1, $2)
RETURNING id, name, created_at, updated_at, parent_id
`

type CreateProductCategoryParams struct {
	Name     string
	ParentID sql.NullInt64
}

func (q *Queries) CreateProductCategory(ctx context.Context, arg CreateProductCategoryParams) (ProductCategory, error) {
	row := q.db.QueryRow(ctx, createProductCategory, arg.Name, arg.ParentID)
	var i ProductCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ParentID,
	)
	return i, err
}

const deleteProductCategory = `-- name: DeleteProductCategory :exec
DELETE
FROM product_categories
WHERE id = $1
`

func (q *Queries) DeleteProductCategory(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteProductCategory, id)
	return err
}

const getProductCategory = `-- name: GetProductCategory :one
SELECT id, name, created_at, updated_at, parent_id
FROM product_categories
WHERE id = $1
`

func (q *Queries) GetProductCategory(ctx context.Context, id int64) (ProductCategory, error) {
	row := q.db.QueryRow(ctx, getProductCategory, id)
	var i ProductCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ParentID,
	)
	return i, err
}

const listProductCategories = `-- name: ListProductCategories :many
SELECT id, name, created_at, updated_at, parent_id
FROM product_categories
ORDER BY name
`

func (q *Queries) ListProductCategories(ctx context.Context) ([]ProductCategory, error) {
	rows, err := q.db.Query(ctx, listProductCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProductCategory
	for rows.Next() {
		var i ProductCategory
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ParentID,
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
