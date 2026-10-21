package cmd

import (
	"ecommerece/config"
	"ecommerece/rest"
)

func Serve() {
	cnf := config.GetConfig()

	rest.Start(cnf)

}
