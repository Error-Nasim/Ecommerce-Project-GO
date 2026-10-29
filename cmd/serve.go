package cmd

import (
	"ecommerece/config"
	"ecommerece/infra/db"
	"ecommerece/product"
	"ecommerece/repo"
	"ecommerece/rest"
	prdctHandler "ecommerece/rest/handlers/product"
	usrHandler "ecommerece/rest/handlers/user"
	middleware "ecommerece/rest/middlewares"
	"ecommerece/user"
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

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//repors
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	//domains
	usrSvc := user.NewService(userRepo)
	prdSvc := product.NewService(productRepo)

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := prdctHandler.NewHandler(middlewares, prdSvc)
	userHandler := usrHandler.NewHandler(cnf, usrSvc)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
