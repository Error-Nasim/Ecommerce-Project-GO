package review

import (
	middleware "ecommerece/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(
	mux *http.ServeMux,
	manager *middleware.Manager,
) {

	mux.Handle(
		"GET /reviews",
		manager.With(
			http.HandlerFunc(h.GetReviews),
		))
}
