-- name: GetPurchaseOrderWithItems :many
SELECT *
FROM purchase_orders as po
         INNER JOIN purchase_order_items AS poi
                    ON po.id == poi.purchase_order_id
WHERE po.id = $1;

-- name: ListPurchaseOrdersWithItems :many
SELECT *
FROM purchase_orders as po
         INNER JOIN purchase_order_items AS poi
                    ON po.id == poi.purchase_order_id
ORDER BY po.id;

-- name: CreatePurchaseOrder :one
INSERT INTO purchase_orders (supplier_id, additional_discount)
VALUES ($1, $2)
RETURNING *;

-- name: CreatePurchaseOrderItems :copyfrom
INSERT INTO purchase_order_items (purchase_order_id, product_id, quantity, amount, discount)
VALUES ($1, $2, $3, $4, $5);