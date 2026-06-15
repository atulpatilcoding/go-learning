package main

import "fmt"

//in go enums don't exist but we can use constants and custom types to achieve similar functionality
type OrderStatus int //We made here a custom type called OrderStatus which is of type int,
// this will help us to create a set of related constants that represent the different states of an order
type PizzaSize string

const (
	Received OrderStatus = iota // iota is a special identifier in Go that is used to create a sequence of related constants,
	//  iota starts at 0 and increments by 1 for each constant in the block
	confirmed
	Prepared
	Delivered
)

const (
	Small  PizzaSize = "small" // Here we are defining constants for pizza sizes using the custom type PizzaSize, which is a string
	Medium PizzaSize = "medium"
	Large  PizzaSize = "large"
)

func changePizzaSize(size PizzaSize) {
	fmt.Println("changing pizza size to", size)
}
func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func main() {
	changeOrderStatus(Prepared)
	changePizzaSize(Small)
}

//  Sr kg Adweta
