package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

// Open close principle
func (p payment) makePayment(amount float32) {
	// razorpayClient := razorpay{}
	// stripeClient := stripe{}
	// razorpayClient.pay(amount)
	// stripeClient.pay(amount)

	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("Making a payment using razorpay. Amount:", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Making a payment using stripe. Amount:", amount)
}

func main() {
	stripePaymentGateway := stripe{}
	// razorpayPaymentGateway := razorpay{}
	newPayment := payment{
		gateway: stripePaymentGateway,
	}
	newPayment.makePayment(100)
}
