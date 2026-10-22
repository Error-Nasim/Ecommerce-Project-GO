package product

import (
	middleware "ecommerece/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(h.GetProducts),
		))
	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(h.CreateProduct),
			h.midlewares.AuthenticateJwt,
		))
	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(h.GetProduct),
		))
	mux.Handle(
		"PUT /products/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateProduct),
			h.midlewares.AuthenticateJwt,
		))
	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteProduct),
			h.midlewares.AuthenticateJwt,
		))
}
