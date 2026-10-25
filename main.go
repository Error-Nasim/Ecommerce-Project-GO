package main

import (
	"fmt"
	"os"
)

// import "ecommerece/cmd"

type People interface {
	PrintDetails()
	RecieveMoney(amount float64) float64
}

type BankUser interface {
	WithdrawMoney(amount float64) float64
}

type user struct {
	Name  string
	Age   int
	Money float64
}

func (obj user) PrintDetails() {
	fmt.Println("Name:", obj.Name)
	fmt.Println("Age:", obj.Age)
	fmt.Println("Money:", obj.Money)
}

func (obj user) RecieveMoney(amount float64) float64 {
	return obj.Money + amount
}

func (obj user) WithdrawMoney(amount float64) float64 {
	return obj.Money - amount
}

func main() {

	// cmd.Serve()
	var usr1 People
	usr1 = user{
		Name:  "Abu Nasim",
		Age:   23,
		Money: 0.50,
	}
	var usr2 BankUser
	usr2 = user{
		Name:  "Sadie Sink",
		Age:   21,
		Money: 4000000000,
	}

	obj, ok := usr2.(user) // convert interface int o orginal struct
	if !ok {
		fmt.Println("Sorry usr2 is not ype of user struct")
		os.Exit(1)
	}
	obj.PrintDetails()
	usr1.PrintDetails()
	usr1.RecieveMoney(100)
	usr2.WithdrawMoney(20000)
}
