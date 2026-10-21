package handlers

import (
	"ecommerece/database"
	"ecommerece/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.List(), 200)
}
