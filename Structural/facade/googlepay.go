package facade

import "fmt"

type GooglePayFacade struct{}

func (p *GooglePayFacade) Pay(amount int) {
	fmt.Println("Amount Payed using GooglePay: ", amount)
}
