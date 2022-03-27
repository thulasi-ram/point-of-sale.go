-- name: GetProduct :one
SELECT *
FROM products
WHERE id = $1;

-- name: ListProducts :many
SELECT *
FROM products
ORDER BY name;

-- name: CreateProduct :one
INSERT INTO products (name, description, category_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteProduct :exec
DELETE
FROM products
WHERE id = $1;