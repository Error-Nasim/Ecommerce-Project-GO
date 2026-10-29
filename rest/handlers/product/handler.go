package product

import (
	middleware "ecommerece/rest/middlewares"
)

type Handler struct {
	midlewares *middleware.Middlewares
	svc        Service
}

func NewHandler(
	midlewares *middleware.Middlewares,
	svc Service,
) *Handler {
	return &Handler{
		midlewares: midlewares,
		svc:        svc,
	}
}
