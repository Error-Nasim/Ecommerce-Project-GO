package product

import (
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

	product, err := h.productRepo.Get(pID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if product == nil {
		util.SendError(w, http.StatusNotFound, "Product Not Found")
		return
	}

	util.SendData(w, product, http.StatusCreated)

}
