package cmd

import (
	"ecommerece/config"
	"ecommerece/repo"
	"ecommerece/rest"
	"ecommerece/rest/handlers/product"
	"ecommerece/rest/handlers/user"
	middleware "ecommerece/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
