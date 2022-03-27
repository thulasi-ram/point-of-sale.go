package products

import (
	"database/sql"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4"
	"net/http"
	"point-of-sale.go/v1/internal/web"
	"point-of-sale.go/v1/products/repository"
	"strconv"
)

func GetProductEndpoint(app *web.App, r *http.Request) *web.APIResponse {
	repo := repository.New(app.DB())

	productID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return web.NewErrorAPIResponse(err, 500)
	}

	product, err := repo.GetProduct(r.Context(), int64(productID))

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return web.NewErrorAPIResponse(err, 404)
		}
		return web.NewErrorAPIResponse(err, 500)
	}

	return web.NewAPIResponse(product, 200)
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

func CreateProductEndpoint(app *web.App, r *http.Request) *web.APIResponse {
	repo := repository.New(app.DB())
	product, err := repo.CreateProduct(r.Context(), repository.CreateProductParams{
		Name:     "Test",
		Category: sql.NullString{String: "Test Category", Valid: true},
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

func RoutesInit(app *web.App) {
	r := app.Router()
	handler := web.HandlerFunc(app)

	r.Route("/api/products", func(r chi.Router) {
		r.Get("/", handler(ListProductsEndpoint))
		r.Post("/", handler(CreateProductEndpoint))
		r.Get("/{id}", handler(GetProductEndpoint))
		r.Delete("/{id}", handler(DeleteProductEndpoint))
	})
}
