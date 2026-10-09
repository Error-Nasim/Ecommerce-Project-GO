package handlers

import (
	"ecommerece/data"
	"ecommerece/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct data.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid json", 400)
		return
	}

	newProduct.ID = len(data.ProductList) + 1
	data.ProductList = append(data.ProductList, newProduct)

	util.SendData(w, newProduct, 201)
}
