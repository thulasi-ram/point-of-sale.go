package controllers

import (
	"github.com/go-chi/chi/v5"
	"point-of-sale.go/v1/internal/web"
)

func InitRoutes(app *web.App) {
	r := app.Router()
	handler := web.HandlerFunc(app)

	r.Route("/api/suppliers", func(r chi.Router) {
		r.Get("/", handler(ListSuppliersEndpoint))
		r.Post("/", handler(CreateSupplierEndpoint))
		r.Get("/{id}", handler(GetSupplierEndpoint))
		r.Delete("/{id}", handler(DeleteSupplierEndpoint))
	})

	r.Route("/api/purchase-orders", func(r chi.Router) {
		r.Post("/", handler(CreatePurchaseOrderEndpoint))
	})
}
