-- name: GetSupplier :one
SELECT *
FROM suppliers
WHERE id = $1;

-- name: ListSuppliers :many
SELECT *
FROM suppliers
ORDER BY name;

-- name: CreateSupplier :one
INSERT INTO suppliers (name, phone, address)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteSupplier :exec
DELETE
FROM suppliers
WHERE id = $1;