package product

import middleware "ecommerece/rest/middlewares"

type Handler struct {
	midlewares *middleware.Middlewares
}

func NewHandler(midlewares *middleware.Middlewares) *Handler {
	return &Handler{
		midlewares: midlewares,
	}
}
