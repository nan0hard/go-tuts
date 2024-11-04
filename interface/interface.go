package main

import "fmt"

type payementer interface {
	pay(amount float32)
}

type payment struct {
	gateway payementer
}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGW := razorpay{}
	// razorpayPaymentGW.pay(amount)

	// stripePaymentGW := stripe{}
	// stripePaymentGW.pay(32)
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to makae a payment
	fmt.Println("Making a payment using razorpay : ", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Making a payment using stripe ", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("Making payment using paypal", amount)
}

func main() {
	// newPayment := payment{}
	// newPayment.makePayment(20)

	paypalGW := paypal{}
	newPayment := payment{
		gateway: paypalGW,
	}

	newPayment.gateway.pay(50)
}
