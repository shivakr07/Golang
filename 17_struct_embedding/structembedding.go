package main

import (
	"fmt"
	"time"
)

/* as we use
composition [code reusability]
and inheritance in oops languages
[all the methods defined for customer will be avaible for customer of order also]
*/

type customer struct {
	name  string
	phone string
}

// if you want to reference customer struct with order struct

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}

func main() {

	// to use

	newOrder := order{
		id:     "1",
		amount: 100,
		status: "received",
	}

	fmt.Println(newOrder)
	// till now you will get customer struct {"" ""}empty value

	//to set the customer values
	// newCustomer := customer{
	// 	name:  "john",
	// 	phone: "1234567890",
	// }

	newOrder2 := order{
		id:     "1",
		amount: 100,
		status: "received",
		// customer: newCustomer,
		// or you can do inline as well instead of defining above
		customer: customer{
			name:  "alan",
			phone: "9999999999",
		},
	}

	fmt.Println(newOrder2)

	//how to change the value of name of newOrder2
	newOrder2.customer.name = "evans"
	fmt.Println(newOrder2)
}
