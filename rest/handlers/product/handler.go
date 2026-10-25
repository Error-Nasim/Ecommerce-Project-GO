package product

import (
	"ecommerece/repo"
	middleware "ecommerece/rest/middlewares"
)

type Handler struct {
	midlewares  *middleware.Middlewares
	productRepo repo.ProductRepo
}

func NewHandler(
	midlewares *middleware.Middlewares,
	productRepo repo.ProductRepo,
) *Handler {
	return &Handler{
		midlewares:  midlewares,
		productRepo: productRepo,
	}
}
