package product

import (
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

	err = h.svc.Delete(pID)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	util.SendData(w, "Sucessfully deleted product", http.StatusCreated)

}
