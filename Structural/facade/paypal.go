package facade

import "fmt"

type PaypalFacade struct{}

func (p *PaypalFacade) Pay(amount int) {
	fmt.Println("Amount Payed using Paypal: ", amount)
}
