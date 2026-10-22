package product

import (
	"ecommerece/database"
	"ecommerece/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid json", 400)
		return
	}

	createdProduct := database.Store(newProduct)

	util.SendData(w, createdProduct, 201)
}
