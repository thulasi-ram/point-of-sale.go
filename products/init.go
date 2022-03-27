package products

import (
	"github.com/go-chi/chi/v5"
	"point-of-sale.go/v1/internal/web"
	"point-of-sale.go/v1/products/controllers"
)

func InitRoutes(app *web.App) {
	r := app.Router()
	handler := web.HandlerFunc(app)

	r.Route("/api/products", func(r chi.Router) {
		r.Get("/", handler(controllers.ListProductsEndpoint))
		r.Post("/", handler(controllers.CreateProductEndpoint))
		r.Get("/{id}", handler(controllers.GetProductEndpoint))
		r.Delete("/{id}", handler(controllers.DeleteProductEndpoint))
	})
}
