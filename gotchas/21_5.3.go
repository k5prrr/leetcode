// Задача найти недочёты и поправить код
package main

import (
	"errors"
	"fmt"
)

type PaymentProcessor interface {
	Process(amount float64) error
	Verify(amount float64) bool
}

type CreditCardProcessor struct {
	limit float64
}

func (c *CreditCardProcessor) Process(amount float64) error {
	if amount > c.limit {
		return errors.New("Limit!")
	}
	fmt.Printf("%.2f using\n")
	c.limit -= amount

	return nil
}
func (c CreditCardProcessor) Verify(amount float64) bool {
	return amount <= c.limit
}

type PayPalProcessor struct {
	balance float64
}

func (p *PayPalProcessor) Process(amount float64) error {
	if amount > p.balance {
		return errors.New("Balance!")
	}
	fmt.Printf("%.2f using\n")
	p.balance -= amount

	return nil
}
func (p *PayPalProcessor) Verify(amount float64) bool {
	return amount <= p.balance
}

func ExecutePayment(processor PaymentProcessor, amount float64) {
	if processor.Verify(amount) {
		err := processor.Process(amount)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	} else {
		fmt.Println("Verif failed")
	}
}

func main() {
	/*// Было
	creditCard := CreditCardProcessor{limit: 100.0}
	payPal := PayPalProcessor{balance: 200.0}

	ExecutePayment(creditCard, 50.0)
	ExecutePayment(&creditCard, 50.0)

	ExecutePayment(&payPal, 150.0)
	ExecutePayment(payPal, 150.0)
	//*/

	creditCard := &CreditCardProcessor{limit: 100.0}
	payPal := &PayPalProcessor{balance: 200.0}

	ExecutePayment(creditCard, 50.0)
	ExecutePayment(creditCard, 50.0)

	ExecutePayment(payPal, 150.0)
	ExecutePayment(payPal, 150.0)
	ExecutePayment(payPal, 150.0)

}
