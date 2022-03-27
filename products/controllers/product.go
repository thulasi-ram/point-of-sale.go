package controllers

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4"
	"net/http"
	"point-of-sale.go/v1/internal/web"
	"point-of-sale.go/v1/products/repository"
	"strconv"
)

type getProductRequest struct {
	Id int64 `json:"id,string" validate:"required"`
}

func GetProductEndpoint(app *web.App, r *http.Request) *web.APIResponse {
	data := &getProductRequest{}
	err := app.Validate(data, r, false, false, true, false)
	if err != nil {
		return web.NewErrorAPIResponse(fmt.Errorf("invalid payload %w", err), 400)
	}

	repo := repository.New(app.DB())

	if err != nil {
		return web.NewErrorAPIResponse(err, 500)
	}

	product, err := repo.GetProduct(r.Context(), data.Id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return web.NewErrorAPIResponse(err, 404)
		}
		return web.NewErrorAPIResponse(err, 500)
	}

	return web.NewAPIResponse(product, 200)
}

type createProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	CategoryID  int64  `json:"category_id" validate:"required"`
}

func CreateProductEndpoint(app *web.App, r *http.Request) *web.APIResponse {
	data := &createProductRequest{}
	err := app.Validate(data, r, true, false, true, false)
	if err != nil {
		return web.NewErrorAPIResponse(fmt.Errorf("invalid payload %w", err), 400)
	}

	repo := repository.New(app.DB())
	product, err := repo.CreateProduct(r.Context(), repository.CreateProductParams{
		Name:        data.Name,
		Description: data.Description,
		CategoryID:  data.CategoryID,
	})

	if err != nil {
		return web.NewErrorAPIResponse(err, 500)
	}

	return web.NewAPIResponse(product, 201)
}

func ListProductsEndpoint(app *web.App, r *http.Request) *web.APIResponse {
	repo := repository.New(app.DB())
	products, err := repo.ListProducts(r.Context())

	if err != nil {
		return web.NewErrorAPIResponse(err, 500)
	}

	return web.NewAPIResponse(products, 201)
}

func DeleteProductEndpoint(app *web.App, r *http.Request) *web.APIResponse {
	repo := repository.New(app.DB())

	productID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return web.NewErrorAPIResponse(err, 500)
	}

	err = repo.DeleteProduct(r.Context(), int64(productID))

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return web.NewErrorAPIResponse(err, 404)
		}
		return web.NewErrorAPIResponse(err, 500)
	}

	return web.NewAPIResponse(map[string]interface{}{
		"product_id": productID,
	}, 200)

}
