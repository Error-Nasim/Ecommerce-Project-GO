package product

import (
	"ecommerece/database"
	"ecommerece/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.List(), 200)
}
