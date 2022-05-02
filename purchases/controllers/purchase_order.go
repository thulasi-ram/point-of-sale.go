package controllers

import (
	"fmt"
	"github.com/shopspring/decimal"
	"net/http"
	"point-of-sale.go/v1/internal/db/repository"
	"point-of-sale.go/v1/internal/types"
	"point-of-sale.go/v1/internal/web"
)

type createPurchaseOrderRequest struct {
	SupplierID         types.ID                         `json:"supplier_id" validate:"required"`
	AdditionalDiscount decimal.Decimal                  `json:"additional_discount" validate:"required"`
	Items              []createPurchaseOrderItemRequest `json:"items" validate:"required"`
}

type createPurchaseOrderItemRequest struct {
	ProductID types.ID        `json:"product_id" validate:"required"`
	Quantity  decimal.Decimal `json:"quantity" validate:"required"`
	Amount    decimal.Decimal `json:"amount" validate:"required"`
	Discount  decimal.Decimal `json:"discount" validate:"required"`
}

func CreatePurchaseOrderEndpoint(app *web.App, r *http.Request) *web.APIResponse {
	data := &createPurchaseOrderRequest{}
	err := app.Validate(data, r, true, false, false, false)
	if err != nil {
		return web.NewErrorAPIResponse(fmt.Errorf("invalid payload %w", err), 400)
	}

	repo := repository.New(app.DB())
	po, err := repo.CreatePurchaseOrder(r.Context(), repository.CreatePurchaseOrderParams{
		SupplierID:         data.SupplierID.Int64(),
		AdditionalDiscount: data.AdditionalDiscount,
	})

	var poReqItems []repository.CreatePurchaseOrderItemsParams
	for _, poi := range data.Items {
		poReqItems = append(poReqItems, repository.CreatePurchaseOrderItemsParams{
			PurchaseOrderID: po.ID,
			ProductID:       poi.ProductID.Int64(),
			Quantity:        poi.Quantity,
			Amount:          poi.Amount,
			Discount:        poi.Discount,
		})

	}

	poItems, err := repo.CreatePurchaseOrderItems(r.Context(), poReqItems)

	if err != nil {
		return web.NewErrorAPIResponse(err, 500)
	}

	return web.NewAPIResponse(map[string]interface{}{
		"purchase_order": po,
		"items":          poItems,
	}, 201)
}
