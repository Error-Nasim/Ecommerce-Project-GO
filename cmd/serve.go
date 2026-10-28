package cmd

import (
	"ecommerece/config"
	"ecommerece/infra/db"
	"ecommerece/repo"
	"ecommerece/rest"
	"ecommerece/rest/handlers/product"
	"ecommerece/rest/handlers/user"
	middleware "ecommerece/rest/middlewares"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

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
