package controllers

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"net/http"
	"point-of-sale.go/v1/internal/db/repository"
	"point-of-sale.go/v1/internal/types"
	"point-of-sale.go/v1/internal/web"
)

type getCustomerRequest struct {
	Id types.ID `json:"id,string" validate:"required"  msgpack:"id"`
}

func GetCustomerEndpoint(app *web.App, r *http.Request) *web.APIResponse {
	data := &getCustomerRequest{}
	err := app.Validate(data, r, false, false, true, false)
	if err != nil {
		return web.NewErrorAPIResponse(fmt.Errorf("invalid payload %w", err), 400)
	}

	repo := repository.New(app.DB())

	if err != nil {
		return web.NewErrorAPIResponse(err, 500)
	}

	customer, err := repo.GetCustomer(r.Context(), data.Id.Int64())

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return web.NewErrorAPIResponse(err, 404)
		}
		return web.NewErrorAPIResponse(err, 500)
	}

	return web.NewAPIResponse(customer, 200)
}

type createCustomerRequest struct {
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Address string `json:"address" validate:"required"`
}

func CreateCustomerEndpoint(app *web.App, r *http.Request) *web.APIResponse {
	data := &createCustomerRequest{}
	err := app.Validate(data, r, true, false, true, false)
	if err != nil {
		return web.NewErrorAPIResponse(fmt.Errorf("invalid payload %w", err), 400)
	}

	repo := repository.New(app.DB())
	customer, err := repo.CreateCustomer(r.Context(), repository.CreateCustomerParams{
		Name:    data.Name,
		Phone:   data.Phone,
		Address: data.Address,
	})

	if err != nil {
		return web.NewErrorAPIResponse(err, 500)
	}

	return web.NewAPIResponse(customer, 201)
}

func ListCustomersEndpoint(app *web.App, r *http.Request) *web.APIResponse {
	repo := repository.New(app.DB())
	customers, err := repo.ListCustomers(r.Context())

	if err != nil {
		return web.NewErrorAPIResponse(err, 500)
	}

	return web.NewAPIResponse(customers, 200)
}

type deleteCustomerRequest struct {
	Id int64 `json:"id,string" validate:"required"`
}

func DeleteCustomerEndpoint(app *web.App, r *http.Request) *web.APIResponse {
	data := &deleteCustomerRequest{}
	err := app.Validate(data, r, false, false, true, false)
	if err != nil {
		return web.NewErrorAPIResponse(fmt.Errorf("invalid payload %w", err), 400)
	}

	repo := repository.New(app.DB())

	err = repo.DeleteCustomer(r.Context(), data.Id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return web.NewErrorAPIResponse(err, 404)
		}
		return web.NewErrorAPIResponse(err, 500)
	}

	return web.NewAPIResponse(map[string]interface{}{
		"customer_id": data.Id,
	}, 200)

}
