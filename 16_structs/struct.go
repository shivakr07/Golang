package main

import (
	"fmt"
	"time"
)

//custom datastructures like classes
//as go don't have classes

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanoseconds precision
}

// there is a constructor kind of thing in oops to setup initial values
// we use a hack : we make a function and generally we keep it with struct definition
// we use new keyword [a convention : anyname]
// till now we have used small letter but in go we use first letter capital [there is a reason]
// and here also not mandatory to assign all the values
// and we return pointer of struct not direct struct
func newOrder(id string, amount float32, status string) *order {
	//initial setup goes here...
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

// to get the same functionality of classes by defining methods
// to connect this method to the struct we use reciver type and below technique
// receiver type o order  [we use first letter of the struct as a convention]
// we also need to use reference to change the value of instance
//and here we don't need to dereferencing [*o.status] as struct auto do this thing

func (o *order) changeStatus(status string) {
	o.status = status
}

// here we don't need * because here we are not changing anything in the value
func (o order) getAmount() float32 {
	return o.amount
}

func main() {

	//to create the instance of the struct order
	myOrder := order{
		id:     "1",
		amount: 100.00,
		status: "received",
	}
	//* it's not mandatory to use all the fields
	//*if you don't assign then it will be zeroed value
	// int=>0, float=>0 string=>"" boolean=>false
	fmt.Println(myOrder)

	myOrder.createdAt = time.Now()
	fmt.Println(myOrder)
	fmt.Println(myOrder.status)

	myOrder2 := order{
		id:        "2",
		amount:    50,
		status:    "delivered",
		createdAt: time.Now(),
	}

	fmt.Println(myOrder2)
	//using method
	myOrder2.changeStatus("confirmed")
	fmt.Println(myOrder2)
	fmt.Println(myOrder2.getAmount())

	//to use constructor
	myOrder3 := newOrder("3", 20, "received")
	fmt.Println(*myOrder3)
	//for individual value need tot to dereference
	fmt.Println(myOrder3.amount)
}

/**
- in go there is a convention that
-> when you want to change the value then use * in struct
-> else you don't need to use the *
*/

// ----------------------------------------
// sometimes we don't use to create the multiple instance of the struct we just want to use that once like doing configuration  or the objects we make in js using literal syntax {}

// we can do the same go to
// you don't need to name the struct

// package main

// import "fmt"

// func main() {

// 	language := struct {
// 		name   string
// 		idGood bool
// 	}{"golang", true}

// 	fmt.Println(language)
// }
