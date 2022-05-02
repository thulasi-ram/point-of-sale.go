package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"point-of-sale.go/v1/internal/db/repository"
	"point-of-sale.go/v1/internal/test"
	"point-of-sale.go/v1/internal/web"
	"testing"
)

func testPurchaseOrderServer(app *web.App, r *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	InitRoutes(app)
	app.Router().ServeHTTP(recorder, r)
	return recorder
}

func TestPurchaseOrderEndpoint(t *testing.T) {
	t.Run("test create purchase order", func(t *testing.T) {
		app := test.SetupTest()

		repo := repository.New(app.DB())
		supplier, err := repo.CreateSupplier(context.Background(), repository.CreateSupplierParams{
			Name:    "TestSupplier",
			Phone:   "xxxxxxxx",
			Address: "test address 1",
		})
		require.Nil(t, err)
		require.NotNil(t, supplier)

		category, err := repo.CreateProductCategory(context.Background(), repository.CreateProductCategoryParams{
			Name: "TestCategory",
		})
		require.Nil(t, err)
		require.NotNil(t, category)

		product, err := repo.CreateProduct(context.Background(), repository.CreateProductParams{
			Name:        "TestProduct",
			Description: "xxxxxxxx",
			CategoryID:  category.ID,
		})
		require.Nil(t, err)
		require.NotNil(t, product)

		payload := map[string]interface{}{
			"supplier_id": supplier.ID,
			"items": []map[string]interface{}{
				{
					"product_id": product.ID,
					"quantity":   "1.5",
					"amount":     "100",
					"discount":   "10",
				},
			},
			"additional_discount": "10.00",
		}
		data, err := json.Marshal(payload)
		require.Nil(t, err)

		req, err := http.NewRequest("POST", "/api/purchase-orders", bytes.NewReader(data))
		require.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")

		response := testPurchaseOrderServer(app, req)

		assert.Equal(t, http.StatusCreated, response.Code)

		poResponse := struct {
			Data struct {
				SupplierID         int64  `json:"supplier_id"`
				AdditionalDiscount string `json:"additional_discount"`
				Items              []struct {
					ProductID string `json:"product_id"`
				}
			} `json:"data"`
		}{}
		poResponse2 := map[string]interface{}{}
		err = json.Unmarshal(response.Body.Bytes(), &poResponse)
		err = json.Unmarshal(response.Body.Bytes(), &poResponse2)
		assert.Nil(t, err)

		//assert.EqualValues(t, supplierResponse.Data, supplier)
		assert.NotEmpty(t, poResponse.Data.SupplierID)

	})
}
