package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"point-of-sale.go/v1/internal/db/repository"
	"point-of-sale.go/v1/internal/test"
	"point-of-sale.go/v1/internal/web"
	"testing"
)

func testSupplierServer(app *web.App, r *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	InitRoutes(app)
	app.Router().ServeHTTP(recorder, r)
	return recorder
}

func TestGetSupplierEndpoint(t *testing.T) {
	t.Run("test get supplier", func(t *testing.T) {

		app := test.SetupTest()

		repo := repository.New(app.DB())
		supplier, err := repo.CreateSupplier(context.Background(), repository.CreateSupplierParams{
			Name:    "TestSupplier",
			Phone:   "xxxxxxxx",
			Address: "test address 1",
		})
		require.Nil(t, err)

		req, err := http.NewRequest("GET", fmt.Sprintf("/api/suppliers/%d", supplier.ID), nil)
		require.Nil(t, err)

		response := testSupplierServer(app, req)

		assert.Equal(t, http.StatusOK, response.Code)

		supplierResponse := struct {
			Data struct {
				ID      int64  `json:"ID"`
				Name    string `json:"Name"`
				Phone   string
				Address string
			} `json:"data"`
		}{}
		supplierResponse2 := map[string]interface{}{}
		err = json.Unmarshal(response.Body.Bytes(), &supplierResponse)
		err = json.Unmarshal(response.Body.Bytes(), &supplierResponse2)
		assert.Nil(t, err)

		//assert.EqualValues(t, supplierResponse.Data, supplier)
		assert.Equal(t, supplierResponse.Data.ID, supplier.ID)
		assert.Equal(t, supplierResponse.Data.Name, supplier.Name)
		assert.Equal(t, supplierResponse.Data.Phone, supplier.Phone)
		assert.Equal(t, supplierResponse.Data.Address, supplier.Address)

	})
	t.Run("test create supplier", func(t *testing.T) {
		app := test.SetupTest()

		payload := map[string]interface{}{
			"name":    "TestSupplier",
			"phone":   "xxxxx",
			"address": "test address 1",
		}
		data, err := json.Marshal(payload)
		require.Nil(t, err)

		req, err := http.NewRequest("POST", "/api/suppliers/", bytes.NewReader(data))
		require.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")

		response := testSupplierServer(app, req)

		assert.Equal(t, http.StatusCreated, response.Code)

		supplierResponse := struct {
			Data struct {
				ID      int64  `json:"ID"`
				Name    string `json:"Name"`
				Phone   string
				Address string
			} `json:"data"`
		}{}
		supplierResponse2 := map[string]interface{}{}
		err = json.Unmarshal(response.Body.Bytes(), &supplierResponse)
		err = json.Unmarshal(response.Body.Bytes(), &supplierResponse2)
		assert.Nil(t, err)

		//assert.EqualValues(t, supplierResponse.Data, supplier)
		assert.NotEmpty(t, supplierResponse.Data.ID)
		assert.Equal(t, supplierResponse.Data.Name, payload["name"])
		assert.Equal(t, supplierResponse.Data.Phone, payload["phone"])
		assert.Equal(t, supplierResponse.Data.Address, payload["address"])

	})
	t.Run("test list suppliers", func(t *testing.T) {
		app := test.SetupTest()

		repo := repository.New(app.DB())
		supplier, err := repo.CreateSupplier(context.Background(), repository.CreateSupplierParams{
			Name:    "TestSupplier",
			Phone:   "xxxxxxxx",
			Address: "test address 1",
		})

		req, err := http.NewRequest("GET", "/api/suppliers/", nil)
		require.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")

		response := testSupplierServer(app, req)

		assert.Equal(t, http.StatusOK, response.Code)

		supplierResponse := struct {
			Data []struct {
				ID      int64  `json:"ID"`
				Name    string `json:"Name"`
				Phone   string
				Address string
			} `json:"data"`
		}{}
		supplierResponse2 := map[string]interface{}{}
		err = json.Unmarshal(response.Body.Bytes(), &supplierResponse)
		err = json.Unmarshal(response.Body.Bytes(), &supplierResponse2)
		assert.Nil(t, err)
		assert.Equal(t, len(supplierResponse.Data), 1)

		assert.Equal(t, supplierResponse.Data[0].ID, supplier.ID)
		assert.Equal(t, supplierResponse.Data[0].Name, supplier.Name)
		assert.Equal(t, supplierResponse.Data[0].Phone, supplier.Phone)
		assert.Equal(t, supplierResponse.Data[0].Address, supplier.Address)

	})
}
