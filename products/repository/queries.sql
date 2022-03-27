-- name: GetProduct :one
SELECT *
FROM products
WHERE id = $1;

-- name: ListProducts :many
SELECT *
FROM products
ORDER BY name;

-- name: CreateProduct :execresult
INSERT INTO products (name, category)
VALUES ($1, $2);

-- name: DeleteProduct :exec
DELETE
FROM products
WHERE id = $1;