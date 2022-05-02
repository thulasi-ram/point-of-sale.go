// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: purchase_orders.queries.sql

package repository

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

const createPurchaseOrder = `-- name: CreatePurchaseOrder :one
INSERT INTO purchase_orders (supplier_id, additional_discount)
VALUES ($1, $2)
RETURNING id, created_at, updated_at, supplier_id, additional_discount
`

type CreatePurchaseOrderParams struct {
	SupplierID         int64
	AdditionalDiscount decimal.Decimal
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

type CreatePurchaseOrderItemsParams struct {
	PurchaseOrderID int64
	ProductID       int64
	Quantity        decimal.Decimal
	Amount          decimal.Decimal
	Discount        decimal.Decimal
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
	AdditionalDiscount decimal.Decimal
	ID_2               int64
	CreatedAt_2        time.Time
	UpdatedAt_2        time.Time
	PurchaseOrderID    int64
	ProductID          int64
	Quantity           decimal.Decimal
	Amount             decimal.Decimal
	Discount           decimal.Decimal
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
	AdditionalDiscount decimal.Decimal
	ID_2               int64
	CreatedAt_2        time.Time
	UpdatedAt_2        time.Time
	PurchaseOrderID    int64
	ProductID          int64
	Quantity           decimal.Decimal
	Amount             decimal.Decimal
	Discount           decimal.Decimal
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
