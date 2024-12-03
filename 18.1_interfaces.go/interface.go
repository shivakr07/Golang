package main

import "fmt"

// interfaces are contract
// we use r in the last like convention
// contract : whatever you mention the interface if any struct implements that then it have to define all those methods mentioned in the interface

//pay(amount float32) bool : you can put if it returns boolean

type paymenter interface {
	pay(amount float32)
	//additional function
	refund(amount float32, account string)
}

// previously we were using the concrete implementation of razorpay,stripe, fake and stripe payment gateway in the struct but now we will not use concrete -> we will use interface/general
type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	//logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

// now if you have to mock the payment
type fake struct{}

func (f fake) pay(amount float32) {
	fmt.Println("making payment using fake pay for testing purpose", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)

}

// we need extra method for refund
func (p paypal) refund(amount float32, account string) {
	fmt.Println("we are processing refund for paypal", amount, account)
}

func main() {

	// razorpayPaymentGw := razorpay{}
	//now its very simple to add any gateway
	paypalGw := paypal{}
	// fakeGw := fake{}
	newPayment := payment{
		// gateway: razorpayPaymentGw,
		// gateway: fakeGw,
		gateway: paypalGw,
	}

	newPayment.makePayment(100)

	//in go we don't implement explicitly like in other languages we need to use implements keyword with classes when they implement any interface

	// in godon't to do it explicitly [don't need to mentioin for a struct that it is implementing the interfce] ->
	//in go if any struct have same signature of that interface then implicitly it implements that interface

	// like here you can see that pay method of razorpay struct follow the same signature [name of method, params, ...] of the pay method of paymenter interface

	// above you can see that our fake struct also implementing the paymenter

	// now we are not modifying only adding [following open-closed principle]

	// interfaces and structs are used implicitly

	//dependency inversion [gateway paymenter] in interface we are doing
	//struct not dependent on concrete thing it dependent on interface/abstraction

}
