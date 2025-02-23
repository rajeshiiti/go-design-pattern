package main

import (
	"fmt"
	"log"
)

type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

type PaypalProcessor struct{}

func (pp *PaypalProcessor) ProcessPayment(amount float64) {
	fmt.Println("Processing payment via paypal... amount: ", amount)
}

type StripeProcessor struct{}

func (sp *StripeProcessor) ProcessPayment(amount float64) {
	fmt.Println("Processing payment via stripe... amount: ", amount)
}

func PaymentProcessorFactory(processorType string) PaymentProcessor {
	switch processorType {
	case "paypal":
		return &PaypalProcessor{}
	case "stripe":
		return &StripeProcessor{}
	default:
		log.Panic("Invalid payment type")
	}
	return nil
}

func main() {
	paypal := PaymentProcessorFactory("paypal")
	paypal.ProcessPayment(100.0)

	stripe := PaymentProcessorFactory("stripe")
	stripe.ProcessPayment(110.9)
}
