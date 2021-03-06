// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: copyfrom.go

package repository

import (
	"context"
)

// iteratorForCreatePurchaseOrderItems implements pgx.CopyFromSource.
type iteratorForCreatePurchaseOrderItems struct {
	rows                 []CreatePurchaseOrderItemsParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreatePurchaseOrderItems) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreatePurchaseOrderItems) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].PurchaseOrderID,
		r.rows[0].ProductID,
		r.rows[0].Quantity,
		r.rows[0].Amount,
		r.rows[0].Discount,
	}, nil
}

func (r iteratorForCreatePurchaseOrderItems) Err() error {
	return nil
}

func (q *Queries) CreatePurchaseOrderItems(ctx context.Context, arg []CreatePurchaseOrderItemsParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"purchase_order_items"}, []string{"purchase_order_id", "product_id", "quantity", "amount", "discount"}, &iteratorForCreatePurchaseOrderItems{rows: arg})
}
