package handlers

import (
	"ecommerece/data"
	"ecommerece/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, data.ProductList, 200)
}
