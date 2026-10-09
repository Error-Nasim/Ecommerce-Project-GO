package handlers

import (
	"ecommerece/data"
	"ecommerece/util"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productID")

	pID, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Please Give Me a VAlid ID", 400)
		return
	}

	for _, product := range data.ProductList {
		if product.ID == pID {
			util.SendData(w, product, 200)
			return
		}
	}
	util.SendData(w, "data not found", 404)

}
