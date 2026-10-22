package product

import (
	"ecommerece/database"
	"ecommerece/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pID, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Please Give Me a Valid ID", 400)
		return
	}

	database.Delete(pID)
	util.SendData(w, "Sucessfully deleted product", 201)

}
