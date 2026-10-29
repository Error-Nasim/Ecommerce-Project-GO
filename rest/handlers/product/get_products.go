package product

import (
	"ecommerece/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList, err := h.svc.List()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	util.SendData(w, productList, http.StatusCreated)
}
