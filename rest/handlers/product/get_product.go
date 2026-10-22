package product

import (
	"ecommerece/database"
	"ecommerece/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pID, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Please Give Me a Valid ID", 400)
		return
	}

	product := database.Get(pID)
	if product == nil {
		util.SendError(w, 404, "Product Not Found")
		return
	}

	util.SendData(w, product, 200)

}
