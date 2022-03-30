// Code generated by sqlc. DO NOT EDIT.
// source: purchase_orders.queries.sql

package repository

import (
	"context"
	"time"

	"github.com/jackc/pgtype"
)

const createPurchaseOrder = `-- name: CreatePurchaseOrder :one
INSERT INTO purchase_orders (supplier_id, additional_discount)
VALUES ($1, $2)
RETURNING id, created_at, updated_at, supplier_id, additional_discount
`

type CreatePurchaseOrderParams struct {
	SupplierID         int64
	AdditionalDiscount pgtype.Numeric
}

func (q *Queries) CreatePurchaseOrder(ctx context.Context, arg CreatePurchaseOrderParams) (PurchaseOrder, error) {
	row := q.db.QueryRow(ctx, createPurchaseOrder, arg.SupplierID, arg.AdditionalDiscount)
	var i PurchaseOrder
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.SupplierID,
		&i.AdditionalDiscount,
	)
	return i, err
}

const createPurchaseOrderItems = `-- name: CreatePurchaseOrderItems :one
INSERT INTO purchase_order_items (purchase_order_id, product_id, quantity, amount, discount)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at, purchase_order_id, product_id, quantity, amount, discount
`

type CreatePurchaseOrderItemsParams struct {
	PurchaseOrderID int64
	ProductID       int64
	Quantity        int32
	Amount          pgtype.Numeric
	Discount        pgtype.Numeric
}

func (q *Queries) CreatePurchaseOrderItems(ctx context.Context, arg CreatePurchaseOrderItemsParams) (PurchaseOrderItem, error) {
	row := q.db.QueryRow(ctx, createPurchaseOrderItems,
		arg.PurchaseOrderID,
		arg.ProductID,
		arg.Quantity,
		arg.Amount,
		arg.Discount,
	)
	var i PurchaseOrderItem
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PurchaseOrderID,
		&i.ProductID,
		&i.Quantity,
		&i.Amount,
		&i.Discount,
	)
	return i, err
}

const getPurchaseOrderWithItems = `-- name: GetPurchaseOrderWithItems :many
SELECT po.id, po.created_at, po.updated_at, supplier_id, additional_discount, poi.id, poi.created_at, poi.updated_at, purchase_order_id, product_id, quantity, amount, discount
FROM purchase_orders as po
         INNER JOIN purchase_order_items AS poi
                    ON po.id == poi.purchase_order_id
WHERE po.id = $1
`

type GetPurchaseOrderWithItemsRow struct {
	ID                 int64
	CreatedAt          time.Time
	UpdatedAt          time.Time
	SupplierID         int64
	AdditionalDiscount pgtype.Numeric
	ID_2               int64
	CreatedAt_2        time.Time
	UpdatedAt_2        time.Time
	PurchaseOrderID    int64
	ProductID          int64
	Quantity           int32
	Amount             pgtype.Numeric
	Discount           pgtype.Numeric
}

func (q *Queries) GetPurchaseOrderWithItems(ctx context.Context, id int64) ([]GetPurchaseOrderWithItemsRow, error) {
	rows, err := q.db.Query(ctx, getPurchaseOrderWithItems, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPurchaseOrderWithItemsRow
	for rows.Next() {
		var i GetPurchaseOrderWithItemsRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.SupplierID,
			&i.AdditionalDiscount,
			&i.ID_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.PurchaseOrderID,
			&i.ProductID,
			&i.Quantity,
			&i.Amount,
			&i.Discount,
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

const listPurchaseOrdersWithItems = `-- name: ListPurchaseOrdersWithItems :many
SELECT po.id, po.created_at, po.updated_at, supplier_id, additional_discount, poi.id, poi.created_at, poi.updated_at, purchase_order_id, product_id, quantity, amount, discount
FROM purchase_orders as po
         INNER JOIN purchase_order_items AS poi
                    ON po.id == poi.purchase_order_id
ORDER BY po.id
`

type ListPurchaseOrdersWithItemsRow struct {
	ID                 int64
	CreatedAt          time.Time
	UpdatedAt          time.Time
	SupplierID         int64
	AdditionalDiscount pgtype.Numeric
	ID_2               int64
	CreatedAt_2        time.Time
	UpdatedAt_2        time.Time
	PurchaseOrderID    int64
	ProductID          int64
	Quantity           int32
	Amount             pgtype.Numeric
	Discount           pgtype.Numeric
}

func (q *Queries) ListPurchaseOrdersWithItems(ctx context.Context) ([]ListPurchaseOrdersWithItemsRow, error) {
	rows, err := q.db.Query(ctx, listPurchaseOrdersWithItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListPurchaseOrdersWithItemsRow
	for rows.Next() {
		var i ListPurchaseOrdersWithItemsRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.SupplierID,
			&i.AdditionalDiscount,
			&i.ID_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.PurchaseOrderID,
			&i.ProductID,
			&i.Quantity,
			&i.Amount,
			&i.Discount,
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