package sales

import (
	"github.com/go-chi/chi/v5"
	"point-of-sale.go/v1/internal/web"
	"point-of-sale.go/v1/sales/controllers"
)

func InitRoutes(app *web.App) {
	r := app.Router()
	handler := web.HandlerFunc(app)

	r.Route("/api/customers", func(r chi.Router) {
		r.Get("/", handler(controllers.ListCustomersEndpoint))
		r.Post("/", handler(controllers.CreateCustomerEndpoint))
		r.Get("/{id}", handler(controllers.GetCustomerEndpoint))
		r.Delete("/{id}", handler(controllers.DeleteCustomerEndpoint))
	})
}
