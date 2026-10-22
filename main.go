package main

import (
	//"ecommerece/cmd"
	"ecommerece/util"
	"fmt"
)

func main() {

	//	cmd.Serve()
	jwt, err := util.CreateJwt("my-secret", util.Payload{
		Sub:         45,
		FirstName:   "Abu",
		LastName:    "Nasim",
		Email:       "nasim@gmail.com",
		IsShopOwner: false,
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jwt)
}
