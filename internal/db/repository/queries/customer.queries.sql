-- name: GetCustomer :one
SELECT *
FROM customers
WHERE id = $1;

-- name: ListCustomers :many
SELECT *
FROM customers
ORDER BY name;

-- name: CreateCustomer :one
INSERT INTO customers (name, phone, address)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteCustomer :exec
DELETE
FROM customers
WHERE id = $1;