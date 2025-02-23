package main

import "fmt"

type PaymentGateway interface {
	ProcessPayment(amount float64)
}

type StripePaymentGateway struct{}

func (s *StripePaymentGateway) ProcessPayment(amount float64) {
	fmt.Println("Processing the payment via StripePaymentGateway")
}

type RazorPayPaymentGateway struct{}

func (r *RazorPayPaymentGateway) ProcessPayment(amount float64) {
	fmt.Println("Processing the payment via RazorPayPaymentGateway")
}

type PaymentMethod struct {
	gateway PaymentGateway
}

func NewPaymentMethod(gateway PaymentGateway) *PaymentMethod {
	return &PaymentMethod{gateway: gateway}
}

func (pm *PaymentMethod) Pay(amount float64) {
	fmt.Println("Processing the payment amount ", amount)
	pm.gateway.ProcessPayment(amount)
}

type CreditCardPaymentMethod struct {
	PaymentMethod
}

func NewCreditCardPaymentMethod(gateway PaymentGateway) *CreditCardPaymentMethod {
	return &CreditCardPaymentMethod{PaymentMethod{gateway: gateway}}
}

type GooglePay struct {
	PaymentMethod
}

func NewGooglePay(gateway PaymentGateway) *GooglePay {
	return &GooglePay{PaymentMethod{gateway: gateway}}
}

func main() {
	razorPayGw := &RazorPayPaymentGateway{}
	stripeGw := &StripePaymentGateway{}

	creditCard := NewCreditCardPaymentMethod(razorPayGw)
	googlePay := NewGooglePay(stripeGw)

	creditCard.Pay(100.0)
	googlePay.Pay(200.0)
}
