package controllers

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"net/http"
	"point-of-sale.go/v1/internal/db/repository"
	"point-of-sale.go/v1/internal/web"
)

type getSupplierRequest struct {
	Id int64 `json:"id,string" validate:"required"`
}

func GetSupplierEndpoint(app *web.App, r *http.Request) *web.APIResponse {
	data := &getSupplierRequest{}
	err := app.Validate(data, r, false, false, true, false)
	if err != nil {
		return web.NewErrorAPIResponse(fmt.Errorf("invalid payload %w", err), 400)
	}

	repo := repository.New(app.DB())

	if err != nil {
		return web.NewErrorAPIResponse(err, 500)
	}

	supplier, err := repo.GetSupplier(r.Context(), data.Id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return web.NewErrorAPIResponse(err, 404)
		}
		return web.NewErrorAPIResponse(err, 500)
	}

	return web.NewAPIResponse(supplier, 200)
}

type createSupplierRequest struct {
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Address string `json:"address" validate:"required"`
}

func CreateSupplierEndpoint(app *web.App, r *http.Request) *web.APIResponse {
	data := &createSupplierRequest{}
	err := app.Validate(data, r, true, false, false, false)
	if err != nil {
		return web.NewErrorAPIResponse(fmt.Errorf("invalid payload %w", err), 400)
	}

	repo := repository.New(app.DB())
	supplier, err := repo.CreateSupplier(r.Context(), repository.CreateSupplierParams{
		Name:    data.Name,
		Phone:   data.Phone,
		Address: data.Address,
	})

	if err != nil {
		return web.NewErrorAPIResponse(err, 500)
	}

	return web.NewAPIResponse(supplier, 201)
}

func ListSuppliersEndpoint(app *web.App, r *http.Request) *web.APIResponse {
	repo := repository.New(app.DB())
	suppliers, err := repo.ListSuppliers(r.Context())

	if err != nil {
		return web.NewErrorAPIResponse(err, 500)
	}

	return web.NewAPIResponse(suppliers, 200)
}

type deleteSupplierRequest struct {
	Id int64 `json:"id,string" validate:"required"`
}

func DeleteSupplierEndpoint(app *web.App, r *http.Request) *web.APIResponse {
	data := &deleteSupplierRequest{}
	err := app.Validate(data, r, false, false, true, false)
	if err != nil {
		return web.NewErrorAPIResponse(fmt.Errorf("invalid payload %w", err), 400)
	}

	repo := repository.New(app.DB())

	err = repo.DeleteSupplier(r.Context(), data.Id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return web.NewErrorAPIResponse(err, 404)
		}
		return web.NewErrorAPIResponse(err, 500)
	}

	return web.NewAPIResponse(map[string]interface{}{
		"supplier_id": data.Id,
	}, 200)

}
