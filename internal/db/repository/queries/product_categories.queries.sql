-- name: GetProductCategory :one
SELECT *
FROM product_categories
WHERE id = $1;

-- name: ListProductCategories :many
SELECT *
FROM product_categories
ORDER BY name;

-- name: CreateProductCategory :one
INSERT INTO product_categories (name, parent_id)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteProductCategory :exec
DELETE
FROM product_categories
WHERE id = $1;