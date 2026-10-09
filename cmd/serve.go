package cmd

import (
	"ecommerece/handlers"
	"ecommerece/middleware"
	"ecommerece/util"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", middleware.Logger(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", middleware.Logger(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("GET /products/{productID}", middleware.Logger(http.HandlerFunc(handlers.GetProductByID)))

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", util.GlobalRouter(mux))
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
