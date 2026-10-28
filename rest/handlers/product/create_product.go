package product

import (
	"ecommerece/repo"
	"ecommerece/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var req ReqCreateProduct

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid json", 400)
		return
	}

	createdProduct, err := h.productRepo.Create(repo.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	})

	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		fmt.Println(err)
		return
	}

	util.SendData(w, createdProduct, http.StatusCreated)
}
