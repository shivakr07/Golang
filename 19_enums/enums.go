package main

import "fmt"

//we can define our own type in go
// type myType string

// we will implement enums using const
// type OrderStatus int

// const (
// 	Received OrderStatus = iota //it is for unsigned int starts from 0 and further consts will be incremented by 1
// 	Confirmed
// 	Prepared
// 	Delivered
// )

// if you want to keep strings
type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

// enumerated types
// go doesn't have builtin enum support
func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func main() {

	changeOrderStatus(Prepared)
	//you can pass any consts from const group
}
