package purchases

import (
	"github.com/go-chi/chi/v5"
	"point-of-sale.go/v1/internal/web"
	"point-of-sale.go/v1/purchases/controllers"
)

func InitRoutes(app *web.App) {
	r := app.Router()
	handler := web.HandlerFunc(app)

	r.Route("/api/suppliers", func(r chi.Router) {
		r.Get("/", handler(controllers.ListSuppliersEndpoint))
		r.Post("/", handler(controllers.CreateSupplierEndpoint))
		r.Get("/{id}", handler(controllers.GetSupplierEndpoint))
		r.Delete("/{id}", handler(controllers.DeleteSupplierEndpoint))
	})
}
