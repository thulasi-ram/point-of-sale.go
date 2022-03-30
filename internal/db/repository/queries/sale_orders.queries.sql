-- name: GetSaleOrderWithItems :many
SELECT *
FROM sale_orders as so
         INNER JOIN sale_order_items AS soi
                    ON so.id == soi.sale_order_id
WHERE so.id = $1;

-- name: ListSaleOrdersWithItems :many
SELECT *
FROM sale_orders as so
         INNER JOIN sale_order_items AS soi
                    ON so.id == soi.sale_order_id
ORDER BY so.id;

-- name: CreateSaleOrder :one
INSERT INTO sale_orders (customer_id, additional_discount)
VALUES ($1, $2)
RETURNING *;

-- name: CreateSaleOrderItems :one
INSERT INTO sale_order_items (sale_order_id, product_id, quantity, amount, discount)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;