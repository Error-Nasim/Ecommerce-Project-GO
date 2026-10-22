package cmd

import (
	"ecommerece/config"
	"ecommerece/rest"
	"ecommerece/rest/handlers/product"
	"ecommerece/rest/handlers/review"
	"ecommerece/rest/handlers/user"
	middleware "ecommerece/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()
	reviewHandler := review.NewHandler()

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
		reviewHandler,
	)
	server.Start()
}
