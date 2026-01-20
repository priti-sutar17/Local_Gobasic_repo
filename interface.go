package main

import "fmt"

// creating payment gateway

// creating interface
type paymenter interface {
	pay(amount float32)
}

// created struct
type payment struct {
	// gateway stripe
	// gateway razorpay

	gateway paymenter
}

// here we created method
// open close principle
func (p payment) makepayment(amount float32) {
	// razorpayPaymentGW := razorpay{}
	// razorpayPaymentGW.pay(amount)
	// stripePaymentGW := stripe{}
	// stripePaymentGW.pay(amount)
	p.gateway.pay(amount)

}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	//logic to make payment
	fmt.Println("MAking payment using razorpay", amount)

}

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	fmt.Println("MAking payment using Stripe", amount)

// }

type fakePayment struct{}

func (f fakePayment) pay(amount float32) {
	fmt.Println("making payment using fake gateway for testing purpose")
}

type payPal struct{}

func (p payPal) pay(amount float32) {
	fmt.Println("making payment using Paypal gateway", amount)
}

func main() {
	// stripePaymentGW := stripe{}
	// fakeGW := fakePayment{}
	// razorpayPaymentGW := razorpay{}
	payPalGW := payPal{}

	newpayment := payment{
		// gateway: stripePaymentGW,
		// gateway: razorpayPaymentGW,
		// gateway: fakeGW,
		gateway: payPalGW,
	}
	newpayment.makepayment(100)
}
