package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float64)
}

type CreditCard struct {
	Name string
}

func (c *CreditCard) Pay(amount float64) {
	fmt.Printf("Paying using the credit card method, user name: %s and amount %.2f \n", c.Name, amount)
}

type PayPal struct {
	Email string
}

func (pp *PayPal) Pay(amount float64) {
	fmt.Printf("Paying using paypal email: %s , and amount %.2f \n", pp.Email, amount)
}

type Bitcoin struct {
	WalletId string
}

func (b *Bitcoin) Pay(amount float64) {
	fmt.Printf("Paying using the bitcoin, wallet id: %s, amount: %.2f \n", b.WalletId, amount)
}

type ShoppingCard struct {
	paymentMethod PaymentMethod
}

func (sc *ShoppingCard) SetPaymentMethod(paymentMethod PaymentMethod) {
	sc.paymentMethod = paymentMethod
}

func (sc *ShoppingCard) Checkout(amount float64) {
	if sc.paymentMethod == nil {
		fmt.Println("Error: No payment defined")
		return
	}
	sc.paymentMethod.Pay(amount)
}

func main() {

	cart := &ShoppingCard{}

	creditCard := &CreditCard{Name: "Ram"}
	paypal := &PayPal{Email: "abc@gad"}
	bitcoin := &Bitcoin{WalletId: "adfs"}

	cart.SetPaymentMethod(creditCard)
	cart.Checkout(123.12)

	cart.SetPaymentMethod(paypal)
	cart.Checkout(13.2)

	cart.SetPaymentMethod(bitcoin)
	cart.Checkout(12.12)
}
