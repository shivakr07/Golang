// package main

// import "fmt"

// // we will mimic payment gateway [razorpay] kind of thing
// type payment struct{}

// func (p payment) makePayment(amount float32) {
// 	//to use razorpay here we need to make an instance
// 	// razorpayPaymentGw := razorpay{}
// 	// razorpayPaymentGw.pay(amount)

// 	//for stripe
// 	stripePaymentGw := stripe{}
// 	stripePaymentGw.pay(amount)
// }

// type razorpay struct{}

// func (r razorpay) pay(amount float32) {
// 	//logic to make payment
// 	fmt.Println("making payment using razorpay", amount)
// }

// // for stripe
// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	fmt.Println("making payment using stripe", amount)
// }

// func main() {

// 	//to use create instance of payment
// 	newPayment := payment{}
// 	newPayment.makePayment(100)
// }

// till now you see things are working fine
// but problem arises when you want to use stripe instead of razorpay or
// you want to keep razorpay but also want to integrate the stripe

// you see we need to make changes above in the function of makePayment also
// we are required to modify the existing code
//-> we are violating  open-closed[solid principle]
// it should be open for extension and closed for modification

// ** we will enhance this thing one by one
// ---------------------------------------------------------
// 1. we are making new instance of stripe / razorpay on every call so
// ** so whenever we build project in go
// we create the instance of dependency[razorpay, stripe] in the struct itself
// like gatewat stripe in struct of payment

package main

import "fmt"

type payment struct {
	gateway stripe
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe", amount)
}

// mocking of fakepayment to check problem in testing with this
// type fakePayment struct{}

// func (f fakePayment) pay(amount float32) {
// 	fmt.Println("making payment using fake gateway for testing")
// }

func main() {
	// newPayment := payment{}
	// newPayment.makePayment(100)

	// but there is a problem as till now we haven't used the stripe so we are not getting the error
	// but when you will use then you will get the error

	// in real world we will get error
	// because we need to pass the paymentgateway in the payment
	// so we need to pass the instance of stripe here
	// -----------------

	stripePaymentGw := stripe{}
	newPayment := payment{
		gateway: stripePaymentGw,
	}

	newPayment.makePayment(100)

	// previously we have learned that in struct fields are optional but we have to instantiate [constructor pattern]
	// by making a function as we have seen
	// and params of function   while instantiation

	// *still we have problem
	// because assume we want to integrate the razorpay again then we need to do following changes as below
	/*
		- create a new instance of razorpay
		razorpayPaymentGw := razorpay{}
		newPayment := payment{
			gateway : razorpayPaymentGw
		}

		and also in struct
		type payment struct {
			gateway razorpay
		}
	*/
	//problem1. still we need to change in struct
	//problem2. in testing : here we are using the razorpay so while testing also we need to pass the exact/actual razorpay [when you are calling] [instead of dummy payment]
	//? what else solution ?????

	// fakeGw := fakePayment{}
	// newPayment := payment{
	// 	gateway: fakeGw, //but we have to pass stripe
	// }
	// newPayment.gateway.pay(100)

	// so the solution is interface
}
